package main

import (
	"bytes"
	"fmt"
	pb "github.com/alexchen/go_test/proto"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"net/http"
)

func main() {
	MockHandleLog()
}

func MockHandleLog() {
	request := makeRequest()
	req, err := proto.Marshal(request)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	resp, err := http.Post("http://127.0.0.1:8890/personList", "application/x-binary", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &pb.PersonListResponse{}
	proto.Unmarshal(respBytes, respObj)
	fmt.Println(respObj.Women)
}

func makeRequest() *pb.PersonListRequest {
	request := &pb.PersonListRequest{
		Name: "doit",
		Age:  12,
	}
	return request
}
