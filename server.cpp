#include <thread>

#ifdef _MSC_VER
#pragma comment(lib, "ws2_32.lib")
#endif

bool StartRakNet(int port, int broadcastInterval, const char* broadcastMsg);

int main()
{
	const char* msg = "hello msgaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaamsg";

	std::thread thrd1(StartRakNet, 5001, 100, msg);

	for (;;)
	{
		std::this_thread::sleep_for(std::chrono::seconds(1000));
	}
}