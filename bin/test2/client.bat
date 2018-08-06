
set ip=101.132.47.70
set interval=66

cd ..
start gochart.exe --showtext1=%interval%
cd test2

REM fast mode
start kcpclient.exe --ip=%ip% --mode=0 --interval=%interval%

REM fec mode
start kcpclient.exe --ip=%ip% --mode=1 --interval=%interval%
