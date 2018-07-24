#include <thread>

#ifdef _MSC_VER
#pragma comment(lib, "ws2_32.lib")
#endif

bool ConnectRakNetServer(const char*sip, int sport, int updateInterval);

int main()
{
	const int updateInterval = 25;		// 更新间隔
	const char* server = "127.0.0.1";	// 测试服务器地址
	int raknetPort = 5001;				// raknet server 端口号

	std::thread thrd1(ConnectRakNetServer, server, raknetPort, updateInterval);

	for (;;)
	{
		std::this_thread::sleep_for(std::chrono::seconds(1000));
	}
}