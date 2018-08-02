set CURDIR=%~dp0
set GOPATH=%CURDIR%:\..\..\..\..\..\;d:\temp
set GOBIN=%CURDIR%\..\bin
go install kcpclient.go
go install kcpserver.go

pause