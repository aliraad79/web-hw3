package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func getCacheClient() CacheServiceClient {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := NewCacherClient(conn)

	getCacheKey(client, "ali")
	setCacheKey(client, "ali", "test")
	getCacheKey(client, "ali")
	clearCache(client)

	return client
}

func getCacheKey(client CacheServiceClient, key string) {
	response, err := client.GetKey(context.Background(), &Str{Val: key})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("response : ", response.GetVal())
	}
}

func setCacheKey(client CacheServiceClient, key string, value string) {
	response, err := client.SetKey(context.Background(), &Cache{Key: key, Value: value})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

func clearCache(client CacheServiceClient) {
	response, err := client.Clear(context.Background(), &Empty{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
