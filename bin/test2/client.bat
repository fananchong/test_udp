chcp 65001

set ip=101.132.47.70
set interval=33

cd ..
start gochart.exe --showtext1="客户端每隔 33ms 发送 400byte 数据给服务器，服务器回发，到客户端收到包的时间间隔"
cd test2

REM fast mode
start kcpclient.exe --ip=%ip% --mode=0 --interval=%interval%

REM fec mode
start kcpclient.exe --ip=%ip% --mode=1 --interval=%interval%
