FROM golang:latest
MAINTAINER "chentao"
WORKDIR $GOPATH/src/gin-example
ADD . $GOPATH/src/gin-example
#RUN go build main.go
EXPOSE 8000
ENTRYPOINT ["./main"]