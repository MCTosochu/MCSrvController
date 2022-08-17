#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o build/mcsrv_linux main.go
GOOS=windows GOARCH=amd64 go build -o build/mcsrv_win.exe main.go
