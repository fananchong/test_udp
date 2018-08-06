set CURDIR=%~dp0
set GOPATH=%CURDIR%;%CURDIR%\..\..\..\..\..\
set GOBIN=%CURDIR%\..\bin\test2\
go install kcpclient.go common.go
go install kcpserver.go common.go

pause