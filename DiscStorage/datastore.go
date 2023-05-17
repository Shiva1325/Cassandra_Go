package storage

import (
	"Cassandra_Go/pb"
	"encoding/json"
	"log"

	barrel "github.com/mr-karan/barreldb"
)

func InitializeStore(id int) *barrel.Barrel {
	barrel, err := barrel.Init(id, barrel.WithDir("data/"))
	if err != nil {
		log.Printf("Failed to initialize db: %v", err)
	}
	return barrel
}
func SaveUser(user_details *pb.User, key string, diskstorage *barrel.Barrel) *pb.Response {
	byte_data, _ := json.Marshal(user_details)
	diskstorage.Put(key, byte_data)
	return &pb.Response{Status: true}
}
