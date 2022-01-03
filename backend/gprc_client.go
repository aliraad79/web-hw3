package main

import (
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

	client := NewCacheServiceClient(conn)

	return client
}

// func getCacheKey(client CacheServiceClient, key string) {
// 	response, err := client.GetKey(context.Background(), &Str{Val: key})
// 	if err != nil {
// 		log.Fatalf("Error when calling GetKey: %s", err)
// 	}
// 	fmt.Println(response)
// }

// func setCacheKey(client CacheServiceClient, key string, value string) {
// 	response, err := client.SetKey(context.Background(), &Cache{Key: key, Value: value})
// 	if err != nil {
// 		log.Fatalf("Error when calling SetKey: %s", err)
// 	}
// 	fmt.Println(response)
// }
