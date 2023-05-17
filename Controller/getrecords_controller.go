package controller

import (
	"Cassandra_Go/pb"
	"context"
)

func GetData(key *pb.GetUserData, clientservice *CassandraClientService) (*pb.SavedRecords, error) {
	response, _ := clientservice.cassandraclientservice.GetData(context.Background(), key)
	return response, nil
}
