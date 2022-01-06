package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func getCacheClient() CacherClient {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := NewCacherClient(conn)

	return client
}

func getFromCache(client CacherClient, key string) (string, error) {
	response, err := client.Get(context.Background(), &GetBody{Key: key})
	if err != nil {
		fmt.Println(err)
		return response.GetValue(), err
	} else {
		fmt.Println("Cache Get response : ", response.GetValue())
		return response.GetValue(), nil
	}
}

func setInCache(client CacherClient, key string, value string) {
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

func CreateCacheKey(note_id int) string {
	return "note_" + strconv.FormatInt(int64(note_id), 10)
}

func getNote(note_id int, db *gorm.DB, cacheClient CacherClient) (string, error) {
	var note Note

	cacheKey := CreateCacheKey(note_id)
	if result, cacheErr := getFromCache(cacheClient, cacheKey); cacheErr != nil {
		return result, nil
	}

	if err := db.First(&note, note_id); err.Error != nil {
		return "", err.Error
	}
	return "DB Success", nil

}
