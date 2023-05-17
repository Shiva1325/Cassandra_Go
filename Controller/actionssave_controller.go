package controller

import (
	"Cassandra_Go/pb"
	"context"

	"github.com/google/uuid"
)

type CassandraClientService struct {
	cassandraclientservice pb.CassandraServiceClient
}

func NewClientCassandraService(clientservice pb.CassandraServiceClient) *CassandraClientService {
	return &CassandraClientService{
		cassandraclientservice: clientservice,
	}
}

func SaveData(user_details *pb.User, clientservice *CassandraClientService) (*pb.Response, error) {
	user_data := &pb.User{
		Id:      int32(uuid.New().ID()),
		Name:    user_details.Name,
		Country: user_details.Country,
		Hashkey: int64(0),
	}
	response, _ := clientservice.cassandraclientservice.SaveData(context.Background(), user_data)
	return response, nil
}
