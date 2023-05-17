package service

import (
	storage "Cassandra_Go/DiscStorage"
	"Cassandra_Go/MemStorage"
	"Cassandra_Go/pb"
	"context"
	"log"
	"math"
	"strconv"

	barrel "github.com/mr-karan/barreldb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CassandraService struct {
	pb.UnimplementedCassandraServiceServer
	UserStore MemStorage.UserStore
	//Partitioner partitioner
	HostData MemStorage.HostData
	//Ringtoken   ringtoken
	DiskStorage *barrel.Barrel
}

func NewCassandraService(userstore MemStorage.UserStore, hoststore MemStorage.HostData, diskstorage *barrel.Barrel) *CassandraService {
	return &CassandraService{
		UserStore:   userstore,
		HostData:    hoststore,
		DiskStorage: diskstorage,
	}
}

func (cs *CassandraService) SaveData(ctx context.Context, user_details *pb.User) (*pb.Response, error) {
	key := HashFunction([]byte(user_details.Country))
	user_details.Hashkey = key
	response := &pb.Response{
		Status: false,
	}
	hashkey := int64(math.Abs(float64(key)))
	response, _ = cs.UserStore.SaveUserData(user_details)
	host_data := cs.HostData.GetHostId()
	if (hashkey%3)+1 == int64(host_data) {
		storage.SaveUser(user_details, strconv.Itoa(int(hashkey)), cs.DiskStorage)
	} else {
		reply := pb.Response{}
		node_addr := GetServerAddress(int32((hashkey % 3) + 1))
		node_conn, _ := ConnectToPeer(int64((hashkey%3)+1), node_addr)
		err := node_conn.Invoke(context.Background(), "/proto.CassandraService/SaveData", user_details, &reply)
		if err != nil {
			log.Printf("Error: %v", err)
		}
		node_conn.Close()
	}
	return response, nil
}

func (cs *CassandraService) GetData(ctx context.Context, key *pb.GetUserData) (*pb.SavedRecords, error) {
	key.Key = HashFunction([]byte(key.Countryname))
	hashKey := int64(math.Abs(float64(key.Key)))
	response := &pb.SavedRecords{
		Records: make(map[uint32]*pb.User),
	}
	host_data := cs.HostData.GetHostId()
	if (hashKey%3)+1 == int64(host_data) {
		response, _ = cs.UserStore.GetUserData(key)
	} else {
		reply := pb.SavedRecords{}
		node_addr := GetServerAddress(int32((hashKey % 3) + 1))
		node_conn, _ := ConnectToPeer(int64((hashKey%3)+1), node_addr)
		err := node_conn.Invoke(context.Background(), "/proto.CassandraService/GetData", key, &reply)
		if err != nil {
			log.Printf("Error: %v", err)
		}
		node_conn.Close()
		response = &reply
	}
	return response, nil
}

func HashFunction(data []byte) int64 {
	length := len(data)
	var h1, h2, k1, k2 int64
	nBlocks := length / 16
	for i := 0; i < nBlocks; i++ {
		k1, k2 = getBlock(data, i)
		k1 *= c1
		k1 = rotl(k1, 31)
		k1 *= c2
		h1 ^= k1
		h1 = rotl(h1, 27)
		h1 += h2
		h1 = h1*5 + 0x52dce729
		k2 *= c2
		k2 = rotl(k2, 33)
		k2 *= c1
		h2 ^= k2
		h2 = rotl(h2, 31)
		h2 += h1
		h2 = h2*5 + 0x38495ab5
	}
	tail := data[nBlocks*16:]
	k1 = 0
	k2 = 0
	switch length & 15 {
	case 15:
		k2 ^= block(tail[14]) << 48
		fallthrough
	case 14:
		k2 ^= block(tail[13]) << 40
		fallthrough
	case 13:
		k2 ^= block(tail[12]) << 32
		fallthrough
	case 12:
		k2 ^= block(tail[11]) << 24
		fallthrough
	case 11:
		k2 ^= block(tail[10]) << 16
		fallthrough
	case 10:
		k2 ^= block(tail[9]) << 8
		fallthrough
	case 9:
		k2 ^= block(tail[8])
		k2 *= c2
		k2 = rotl(k2, 33)
		k2 *= c1
		h2 ^= k2
		fallthrough
	case 8:
		k1 ^= block(tail[7]) << 56
		fallthrough
	case 7:
		k1 ^= block(tail[6]) << 48
		fallthrough
	case 6:
		k1 ^= block(tail[5]) << 40
		fallthrough
	case 5:
		k1 ^= block(tail[4]) << 32
		fallthrough
	case 4:
		k1 ^= block(tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= block(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= block(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= block(tail[0])
		k1 *= c1
		k1 = rotl(k1, 31)
		k1 *= c2
		h1 ^= k1
	}
	h1 ^= int64(length)
	h2 ^= int64(length)
	h1 += h2
	h2 += h1
	h1 = fmix(h1)
	h2 = fmix(h2)
	h1 += h2
	return h1
}

func GetServerAddress(peerId int32) string {
	BASE_IP := "0.0.0.0"
	port := ":1200" + strconv.Itoa(int(peerId))
	addr := BASE_IP + port
	return addr
}

func ConnectToPeer(peerId int64, addr string) (*grpc.ClientConn, error) {
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientconn, err := grpc.Dial(addr, transportOption)
	if err != nil {
		log.Printf("Cannot Dail server %v", peerId)
	}
	return clientconn, nil
}
