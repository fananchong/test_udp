#include <chrono>

long long get_tick_count(void)
{
	typedef std::chrono::time_point<std::chrono::system_clock, std::chrono::nanoseconds> nanoClock_type;
	nanoClock_type tp = std::chrono::time_point_cast<std::chrono::nanoseconds>(std::chrono::system_clock::now());
	return long long(tp.time_since_epoch().count() / 1000000);
}