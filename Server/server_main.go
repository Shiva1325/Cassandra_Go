package main

import (
	diskstorage "Cassandra_Go/DiscStorage"
	MemStorage "Cassandra_Go/MemStorage"
	service "Cassandra_Go/Service"
	"Cassandra_Go/pb"
	"flag"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	Id := flag.Int("id", 1, "server ID")
	defer func() {
		log.Println("Panic detected")
	}()
	flag.Parse()
	//construct the IP
	serverId := *Id
	serverId_string := strconv.Itoa(serverId)
	IP_BASE := "0.0.0.0"
	port_number := 1200 + serverId
	port := ":1200" + serverId_string
	IP := IP_BASE + port
	//Initialize the Listener and the servers
	listener, err := net.Listen("tcp", IP)
	userstore := MemStorage.NewUserStore()
	diskstore := diskstorage.InitializeStore(serverId)
	hoststore := MemStorage.HostDetailsStore()
	hoststore.SaveHostData(int32(serverId), int32(port_number))
	MemStorage.RestoreDB(diskstore)
	cassandraService := service.NewCassandraService(userstore, hoststore, diskstore)
	if err != nil {
		log.Fatal(err)
	}
	grpcserver := grpc.NewServer()
	pb.RegisterCassandraServiceServer(grpcserver, cassandraService)
	err = grpcserver.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}
