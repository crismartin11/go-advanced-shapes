@echo off
cls
echo :: Building project.

set lambdaName="main"
set directory="build"

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -o %directory%/%lambdaName% cmd/main.go

echo :: Build finished.