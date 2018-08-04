#include <thread>
#include <string.h>

bool StartRakNetServer(int port, int broadcastInterval, const char* broadcastMsg, int broadcastMsgLen);

int main(int argn, char* argv[])
{
	const int dataSzie = 400;			// 数据包大小
	int broadcastInterval = 100;			// 广播间隔
	const int raknetPort = 5001;			// raknet server 端口号

	if (argn > 1)
	{
		broadcastInterval = atoi(argv[1]);
	}

	char msg[dataSzie] = { 0 };
	memset(msg, 97, dataSzie - 1);	// msg =  "aaaaa ... aaaaaa"
	msg[dataSzie - 1] = 0;

	std::thread thrd1(StartRakNetServer, raknetPort, broadcastInterval, msg, dataSzie);

	for (;;)
	{
		std::this_thread::sleep_for(std::chrono::seconds(1000));
	}
}
