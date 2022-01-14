package main

import (
	context "context"
	"encoding/json"
	"fmt"

	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

var cache Cache

type Server struct {
	UnimplementedCacherServer
}

func New() *Server {
	return &Server{}
}

func (s *Server) Set(ctx context.Context, req *SetBody) (*Response, error) {
	cache.set(req.GetKey(), req.GetValue())
	return &Response{Result: "success"}, nil
}

func (s *Server) Get(ctx context.Context, req *GetBody) (*GetResponse, error) {
	value := cache.get(req.GetKey())
	return &GetResponse{Value: value}, nil
}

func (s *Server) Clear(ctx context.Context, req *Empty) (*Response, error) {
	cache.clear()
	return &Response{Result: "success"}, nil
}

type Configuration struct {
	Capacity int
	Port     string
}

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	cache = initCache(configuration.Capacity)
	lis, err := net.Listen("tcp", "0.0.0.0:"+configuration.Port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))

	RegisterCacherServer(grpcServer, New())
	log.Println("Cache started in secure mode on 0.0.0.0:" + configuration.Port)
	grpcServer.Serve(lis)
}
