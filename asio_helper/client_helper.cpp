#include "client_helper.h"

// 这是只是使用 asio 能使用tcp，把数据发给 gochart，非正式实现。

std::shared_ptr<AsioClient> InitAsioClient(const char* ip, const char* port)
{
	static boost::asio::io_context io_context;
	static tcp::resolver resolver(io_context);
	static auto endpoints = resolver.resolve(ip, port);
	auto client = std::make_shared<AsioClient>(io_context, endpoints);
	std::thread t([]() { io_context.run(); });
	return client;
}