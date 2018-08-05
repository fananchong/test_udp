
set ip=101.132.47.70

cd ..
start gochart.exe --showtext1=66
cd test1

REM start tcpclient.exe --ip=%ip%
start kcpclient.exe --ip=%ip%
start x64-Release\client\RelWithDebInfo\client.exe %ip%

