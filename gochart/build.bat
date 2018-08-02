set CURDIR=%~dp0
set GOPATH=%CURDIR%:\..\..\..\..\..\;d:\temp
set GOBIN=%CURDIR%\..\bin
go install ./...

pause