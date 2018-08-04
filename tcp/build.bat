set CURDIR=%~dp0
set GOPATH=%CURDIR%;%CURDIR%\..\..\..\..\..\
set GOBIN=%CURDIR%\..\bin
go install tcpclient.go
go install tcpserver.go

pause
