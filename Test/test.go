package Test

import (
	"Cassandra_Go/pb"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type User struct {
	User []pb.User
}

func GenerateData(id int) *User {
	jsonfile, err := os.Open("Test/Test_Data_" + strconv.Itoa(id) + ".json")
	if err != nil {
		log.Printf("Error reading JSON file: %v", err)
	}
	var user_data User
	bytevalue, _ := ioutil.ReadAll(jsonfile)
	json.Unmarshal(bytevalue, &user_data.User)
	return &user_data
}
