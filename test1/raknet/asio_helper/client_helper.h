#include <boost/asio.hpp>
#include <memory>
#include <deque>

// 这是只是使用 asio 能使用tcp，把数据发给 gochart，非正式实现。

using boost::asio::ip::tcp;

typedef std::string message_type;

typedef std::deque<message_type> message_queue;

class AsioClient
{
public:
	AsioClient(boost::asio::io_context& io_context,
		const tcp::resolver::results_type& endpoints)
		: io_context_(io_context),
		socket_(io_context)
	{
		do_connect(endpoints);
	}

	void Write(const message_type& msg)
	{
		boost::asio::post(io_context_,
			[this, msg]()
		{
			bool write_in_progress = !write_msgs_.empty();
			write_msgs_.push_back(msg);
			if (!write_in_progress)
			{
				do_write();
			}
		});
	}

	void Close()
	{
		boost::asio::post(io_context_, [this]() { socket_.close(); });
	}

private:
	void do_connect(const tcp::resolver::results_type& endpoints)
	{
		boost::asio::async_connect(socket_, endpoints,
			[this](boost::system::error_code ec, tcp::endpoint)
		{
			if (!ec)
			{
				char temp[256];
				boost::asio::async_read(socket_,
					boost::asio::buffer(temp, 256),
					[this](boost::system::error_code ec, std::size_t /*length*/)
				{
					
				});
			}
		});
	}

	void do_write()
	{
		boost::asio::async_write(socket_,
			boost::asio::buffer(write_msgs_.front().data(),
				write_msgs_.front().length()),
			[this](boost::system::error_code ec, std::size_t /*length*/)
		{
			if (!ec)
			{
				write_msgs_.pop_front();
				if (!write_msgs_.empty())
				{
					do_write();
				}
			}
			else
			{
				socket_.close();
			}
		});
	}

private:
	boost::asio::io_context& io_context_;
	tcp::socket socket_;
	message_queue write_msgs_;
};

std::shared_ptr<AsioClient> InitAsioClient(boost::asio::io_context &io_context, const char* ip, const char* port);