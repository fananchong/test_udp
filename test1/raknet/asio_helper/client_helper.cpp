#include "client_helper.h"

// 这是只是使用 asio 能使用tcp，把数据发给 gochart，非正式实现。

std::shared_ptr<AsioClient> InitAsioClient(boost::asio::io_context &io_context, const char* ip, const char* port)
{
	tcp::resolver resolver(io_context);
	auto endpoints = resolver.resolve(ip, port);
	auto client = std::make_shared<AsioClient>(io_context, endpoints);
	return client;
}