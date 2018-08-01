#include <thread>
#include "client_helper.h"

bool ConnectRakNetServer(const char*sip, int sport, int updateInterval, const std::shared_ptr<AsioClient>& asioClient);

int main(int argn, char* argv[])
{
	const int updateInterval = 25;		// 更新间隔
	char* server = "127.0.0.1";			// 测试服务器地址
	int raknetPort = 5001;				// raknet server 端口号

	if (argn > 1)
	{
		server = argv[1];
	}

	boost::asio::io_context io_context;
	std::shared_ptr<AsioClient> asioClient = InitAsioClient(io_context, "127.0.0.1", "3333");
	std::thread thrd1(ConnectRakNetServer, server, raknetPort, updateInterval, asioClient);
	io_context.run();
}