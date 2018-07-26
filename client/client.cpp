#include <thread>
#include "client_helper.h"

bool ConnectRakNetServer(const char*sip, int sport, int updateInterval, const std::shared_ptr<AsioClient>& asioClient);

int main()
{
	const int updateInterval = 25;		// 更新间隔
	const char* server = "127.0.0.1";	// 测试服务器地址
	int raknetPort = 5001;				// raknet server 端口号

	std::shared_ptr<AsioClient> asioClient = InitAsioClient("127.0.0.1", "3333");
	std::thread thrd1(ConnectRakNetServer, server, raknetPort, updateInterval, asioClient);

	for (;;)
	{
		std::this_thread::sleep_for(std::chrono::seconds(1000));
	}
}