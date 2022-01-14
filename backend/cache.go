package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
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

	return NewCacherClient(conn)
}

func getFromCache(key string) (string, error) {
	response, err := cacheClient.Get(context.Background(), &GetBody{Key: key})
	if err != nil || response.GetValue() == "" {
		return response.GetValue(), err
	} else {
		return response.GetValue(), nil
	}
}

func setInCache(key string, value string) {
	_, err := cacheClient.Set(context.Background(), &SetBody{Key: key, Value: value})
	if err != nil {
		fmt.Println(err)
	}
}

func clearCache() {
	response, err := cacheClient.Clear(context.Background(), &Empty{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

func addNoteToCache(note Note) {
	str, _ := json.Marshal(note)
	setInCache("note_"+strconv.FormatInt(int64(note.ID), 10), string(str))
}

func addUserToCache(user User) {
	str, _ := json.Marshal(user)
	setInCache("user_"+user.Username, string(str))
}

func getNote(note_id int) (Note, error) {
	var note Note

	cacheKey := "note_" + strconv.FormatInt(int64(note_id), 10)
	if result, cacheErr := getFromCache(cacheKey); cacheErr == nil && result != "" {
		json.Unmarshal([]byte(result), &note)
		return note, cacheErr
	}

	if err := db.First(&note, note_id); err.Error != nil {
		return Note{}, err.Error
	}

	addNoteToCache(note)
	return note, nil
}

func getUser(user_name string) (User, error) {

	var user User

	cacheKey := "user_" + user_name
	if result, cacheErr := getFromCache(cacheKey); cacheErr == nil && result != "" {
		json.Unmarshal([]byte(result), &user)
		return user, cacheErr
	}

	if err := db.Where("username = ?", user_name).First(&user); err.Error != nil {
		return User{}, err.Error
	}

	addUserToCache(user)
	return user, nil
}
