#!/bin/sh

mkdir ./build -p
GOOS=linux GOARCH=amd64 go build -o build/mcsrv_linux_amd64 main.go
GOOS=linux GOARCH=arm64 go build -o build/mcsrv_linux_arm64 main.go
GOOS=windows GOARCH=amd64 go build -o build/mcsrv_win.exe main.go
