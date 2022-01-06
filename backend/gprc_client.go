package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func getCacheClient() CacherClient {

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

func getCacheKey(client CacherClient, key string) {
	response, err := client.Get(context.Background(), &GetBody{Key: key})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("response : ", response.GetValue())
	}
}

func setCacheKey(client CacherClient, key string, value string) {
	response, err := client.Set(context.Background(), &SetBody{Key: key, Value: value})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

func clearCache(client CacherClient) {
	response, err := client.Clear(context.Background(), &Empty{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
