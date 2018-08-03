set CURDIR=%~dp0
set GOPATH=%CURDIR%:\..\..\..\..\..\;d:\temp
set GOBIN=%CURDIR%\..\bin
go install kcpclient.go common.go
go install kcpserver.go common.go

pause