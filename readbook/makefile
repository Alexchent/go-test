GOPATH:=$(shell go env GOPATH)
JIN_GOPATH:=$(shell pwd)
CGO_ENABLED:=0
GOOS:=linux
GOARCH:=amd64

# 编译Mac 64 位系统（inter&&ARM）
build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o readbook

# 编译Linux 64 位系统
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o readbook

## 编译Windows64位系统可执行文件
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o readbook