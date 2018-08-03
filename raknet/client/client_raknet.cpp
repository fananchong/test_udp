#include "MessageIdentifiers.h"
#include "RakPeerInterface.h"
#include "RakNetStatistics.h"
#include "RakNetTypes.h"
#include "BitStream.h"
#include "PacketLogger.h"
#include <assert.h>
#include <cstdio>
#include <cstring>
#include <stdlib.h>
#include "RakNetTypes.h"
#ifdef _WIN32
#include "WindowsIncludes.h" // Sleep
#else
#include <unistd.h> // usleep
#endif
#include "client_helper.h"
#include <vector>

long long get_tick_count(void);

// We copy this from Multiplayer.cpp to keep things all in one file for this example
unsigned char GetPacketIdentifier(RakNet::Packet *p);

bool ConnectRakNetServer(const char*sip, int sport, int updateInterval, const std::shared_ptr<AsioClient>& asioClient)
{
	RakNet::RakNetStatistics *rss;
	// Pointers to the interfaces of our server and client.
	// Note we can easily have both in the same program
	RakNet::RakPeerInterface *client = RakNet::RakPeerInterface::GetInstance();

	// Holds packets
	RakNet::Packet* p;

	char ip[64], serverPort[30], clientPort[30];
	strcpy(ip, sip);
	sprintf(serverPort, "%d", sport);
	strcpy(clientPort, "0");

	client->AllowConnectionResponseIPMigration(false);

	// Connecting the client is very simple.  0 means we don't care about
	// a connectionValidationInteger, and false for low priority threads
	RakNet::SocketDescriptor socketDescriptor(atoi(clientPort), 0);
	socketDescriptor.socketFamily = AF_INET;
	client->Startup(8, &socketDescriptor, 1);
	client->SetOccasionalPing(true);

	RakNet::ConnectionAttemptResult car = client->Connect(ip, atoi(serverPort), "Rumpelstiltskin", (int)strlen("Rumpelstiltskin"));
	RakAssert(car == RakNet::CONNECTION_ATTEMPT_STARTED);

	printf("\nMy IP addresses:\n");
	unsigned int i;
	for (i = 0; i < client->GetNumberOfAddresses(); i++)
	{
		printf("%i. %s\n", i + 1, client->GetLocalIP(i));
	}

	printf("My GUID is %s\n", client->GetGuidFromSystemAddress(RakNet::UNASSIGNED_SYSTEM_ADDRESS).ToString());

	char message[2048];

	long long preTCPRecvTime = 0;
	std::vector<long long> logs;

	// Loop for input
	while (1)
	{
		// This sleep keeps RakNet responsive
#ifdef _WIN32
		Sleep(1);
#else
		usleep(30 * 1000);
#endif

		// Get a packet from either the server or the client

		for (p = client->Receive(); p; client->DeallocatePacket(p), p = client->Receive())
		{
			// We got a packet, get the identifier with our handy function
			auto packetIdentifier = GetPacketIdentifier(p);

			// Check if this is a network message packet
			switch (packetIdentifier)
			{
			case ID_DISCONNECTION_NOTIFICATION:
				// Connection lost normally
				printf("ID_DISCONNECTION_NOTIFICATION\n");
				break;
			case ID_ALREADY_CONNECTED:
				// Connection lost normally
				printf("ID_ALREADY_CONNECTED with guid %" PRINTF_64_BIT_MODIFIER "u\n", p->guid);
				break;
			case ID_INCOMPATIBLE_PROTOCOL_VERSION:
				printf("ID_INCOMPATIBLE_PROTOCOL_VERSION\n");
				break;
			case ID_REMOTE_DISCONNECTION_NOTIFICATION: // Server telling the clients of another client disconnecting gracefully.  You can manually broadcast this in a peer to peer enviroment if you want.
				printf("ID_REMOTE_DISCONNECTION_NOTIFICATION\n");
				break;
			case ID_REMOTE_CONNECTION_LOST: // Server telling the clients of another client disconnecting forcefully.  You can manually broadcast this in a peer to peer enviroment if you want.
				printf("ID_REMOTE_CONNECTION_LOST\n");
				break;
			case ID_REMOTE_NEW_INCOMING_CONNECTION: // Server telling the clients of another client connecting.  You can manually broadcast this in a peer to peer enviroment if you want.
				printf("ID_REMOTE_NEW_INCOMING_CONNECTION\n");
				break;
			case ID_CONNECTION_BANNED: // Banned from this server
				printf("We are banned from this server.\n");
				break;
			case ID_CONNECTION_ATTEMPT_FAILED:
				printf("Connection attempt failed\n");
				break;
			case ID_NO_FREE_INCOMING_CONNECTIONS:
				// Sorry, the server is full.  I don't do anything here but
				// A real app should tell the user
				printf("ID_NO_FREE_INCOMING_CONNECTIONS\n");
				break;

			case ID_INVALID_PASSWORD:
				printf("ID_INVALID_PASSWORD\n");
				break;

			case ID_CONNECTION_LOST:
				// Couldn't deliver a reliable packet - i.e. the other system was abnormally
				// terminated
				printf("ID_CONNECTION_LOST\n");
				break;

			case ID_CONNECTION_REQUEST_ACCEPTED:
				// This tells the client they have connected
				printf("ID_CONNECTION_REQUEST_ACCEPTED to %s with GUID %s\n", p->systemAddress.ToString(true), p->guid.ToString());
				printf("My external address is %s\n", client->GetExternalID(p->systemAddress).ToString(true));
				break;
			case ID_CONNECTED_PING:
			case ID_UNCONNECTED_PING:
				printf("Ping from %s\n", p->systemAddress.ToString(true));
				break;
			default:
				// It's a client, so just show the message
				//printf("%s\n", p->data);

				if (p->length != 400) {
					printf("data error! p->bitSize = %d", p->bitSize);
					printf("data = %s", p->data);
					exit(0);
				}

				auto now = get_tick_count();
				if (preTCPRecvTime == 0)
				{
					preTCPRecvTime = now;
				}
				auto detal = now - preTCPRecvTime;
				preTCPRecvTime = now;
				char temp[256];
				snprintf(temp, sizeof(temp), "1_%d;", detal);
				asioClient->Write(temp);

				logs.push_back(detal);
				if (logs.size() % 200 == 0) {
					for (size_t i = 0; i < logs.size(); i++)
					{
						printf("%d ", logs[i]);
					}
					printf("\n");
				}

				break;
			}
		}
	}

	// Be nice and let the server know we quit.
	client->Shutdown(300);

	// We're done with the network
	RakNet::RakPeerInterface::DestroyInstance(client);

	return 0;
}

// Copied from Multiplayer.cpp
// If the first byte is ID_TIMESTAMP, then we want the 5th byte
// Otherwise we want the 1st byte
unsigned char GetPacketIdentifier(RakNet::Packet *p)
{
	if (p == 0)
		return 255;

	if ((unsigned char)p->data[0] == ID_TIMESTAMP)
	{
		RakAssert(p->length > sizeof(RakNet::MessageID) + sizeof(RakNet::Time));
		return (unsigned char)p->data[sizeof(RakNet::MessageID) + sizeof(RakNet::Time)];
	}
	else
		return (unsigned char)p->data[0];
}
