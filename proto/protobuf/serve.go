package main

import (
	"fmt"
	pb "go-test/proto"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

func HandleLog(resp http.ResponseWriter, req *http.Request) {
	//contentLength := req.ContentLength
	//fmt.Printf("Content Length Received : %v\n", contentLength)

	request := &pb.PersonListRequest{}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("Error reading", err.Error())
		return
	}
	err = proto.Unmarshal(data, request)
	if err != nil {
		log.Println("Error unmarshaling", err.Error())
		return
	}

	fmt.Println(request.GetName(), request.GetAge())

	womens := make([]*pb.Person, 0)
	womens = append(womens, &pb.Person{
		Id:   1,
		Name: "xiaomi",
		Age:  10,
	})
	result := &pb.PersonListResponse{
		People: womens,
		Count:  1,
	}
	response, err := proto.Marshal(result)
	if err != nil {
		log.Println("Failed to marshal", err.Error())
	}
	resp.Write(response)
}

func main() {
	fmt.Println("Starting the API server...")
	r := mux.NewRouter()
	r.HandleFunc("/personList", HandleLog).Methods("POST")

	server := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8890",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}
	fmt.Println("0.0.0.0:8890 started")

	log.Fatal(server.ListenAndServe())
}
