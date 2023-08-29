CWD=$(shell pwd)
DIR_PROTO=${CWD}/proto

BUILDTIME?=$(shell date +%Y-%m-%dT%T%z)
GITHASH?=$(shell git rev-parse --short HEAD)

.PHONY: default local proto proto-g
default:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

local:
	CGO_ENABLED=0 go build

proto:
	protoc --gogofast_out=plugins=rpcx:${DIR_PROTO} --proto_path ${DIR_PROTO} person.proto

proto-g:
	protoc --gofast_out=${DIR_PROTO} --proto_path ${DIR_PROTO} person.proto