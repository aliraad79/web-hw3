package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func getCacheClient() CacherClient {

	tlsCredentials, tls_err := loadTLSCredentials()
	if tls_err != nil {
		log.Fatal("cannot load TLS credentials: ", tls_err)
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("CACHE_HOST")+":"+os.Getenv("CACHE_PORT"), grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	client := NewCacherClient(conn)

	return client
}

func getFromCache(client CacherClient, key string) (string, error) {
	response, err := client.Get(context.Background(), &GetBody{Key: key})
	if err != nil || response.GetValue() == "" {
		return response.GetValue(), err
	} else {
		return response.GetValue(), nil
	}
}

func setInCache(client CacherClient, key string, value string) {
	_, err := client.Set(context.Background(), &SetBody{Key: key, Value: value})
	if err != nil {
		fmt.Println(err)
	}
}

func clearCache(client CacherClient) {
	response, err := client.Clear(context.Background(), &Empty{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

func addNoteToCache(client CacherClient, note Note) {
	str, _ := json.Marshal(note)
	setInCache(client, "note_"+strconv.FormatInt(int64(note.ID), 10), string(str))
}

func addUserToCache(client CacherClient, user User) {
	str, _ := json.Marshal(user)
	setInCache(client, "user_"+user.Username, string(str))
}

func getNote(note_id int, db *gorm.DB, cacheClient CacherClient) (Note, error) {
	var note Note

	cacheKey := "note_" + strconv.FormatInt(int64(note_id), 10)
	if result, cacheErr := getFromCache(cacheClient, cacheKey); cacheErr == nil && result != "" {
		json.Unmarshal([]byte(result), &note)
		return note, cacheErr
	}

	if err := db.First(&note, note_id); err.Error != nil {
		return Note{}, err.Error
	}

	addNoteToCache(cacheClient, note)
	return note, nil
}

func getUser(user_name string, db *gorm.DB, cacheClient CacherClient) (User, error) {

	var user User

	cacheKey := "user_" + user_name
	if result, cacheErr := getFromCache(cacheClient, cacheKey); cacheErr == nil && result != "" {
		json.Unmarshal([]byte(result), &user)
		return user, cacheErr
	}

	if err := db.Where("username = ?", user_name).First(&user); err.Error != nil {
		return User{}, err.Error
	}

	addUserToCache(cacheClient, user)
	return user, nil
}
