package MemStorage

import (
	"Cassandra_Go/pb"
	"errors"
	"log"
	"net"
	"strconv"
	"sync"

	"github.com/jinzhu/copier"
	barrel "github.com/mr-karan/barreldb"
)

type HostInfo struct {
	peer           []int32
	connectAddress net.IP
	port           int
	hostId         int32
}

// userstore stores the user information on server
type UserStore interface {
	SaveUserData(userdetails *pb.User) (*pb.Response, error)
	GetUserData(key *pb.GetUserData) (*pb.SavedRecords, error)
}
type HostData interface {
	SaveHostData(id int32, port int32) *HostInfo
	GetHostData() *HostInfo
	GetHostId() int32
}

type InMemoryUserStore struct {
	mutex sync.RWMutex
	User  map[int32]*pb.User
}
type InMemoryHostData struct {
	mutex    sync.RWMutex
	HostInfo map[int32]*HostInfo
}

func HostDetailsStore() *InMemoryHostData {
	return &InMemoryHostData{
		HostInfo: make(map[int32]*HostInfo, 0),
	}
}

func NewUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		User: make(map[int32]*pb.User, 0),
	}
}

func RestoreDB(barrel *barrel.Barrel) {

}

func (userstore *InMemoryUserStore) GetUserData(key *pb.GetUserData) (*pb.SavedRecords, error) {
	userstore.mutex.Lock()
	defer userstore.mutex.Unlock()
	response := &pb.SavedRecords{
		Records: make(map[uint32]*pb.User, 0),
	}
	for index, value := range userstore.User {
		if value.Hashkey == key.Key {
			response.Records[uint32(index)] = value
		}
	}
	return response, errors.New("")
}

func (userstore *InMemoryUserStore) SaveUserData(user_details *pb.User) (*pb.Response, error) {
	userstore.mutex.Lock()
	defer userstore.mutex.Unlock()
	response := &pb.Response{
		Status: false,
	}
	user_data := &pb.User{}
	err := copier.Copy(user_data, user_details)
	if err != nil {
		response.Status = false
		log.Printf("Error: %v", err)
	}
	id := user_details.Id
	userstore.User[id] = user_data
	response.Status = true
	return response, errors.New("")
}

func (hd *InMemoryHostData) SaveHostData(id int32, port int32) *HostInfo {
	hd.mutex.Lock()
	defer hd.mutex.Unlock()
	host_data := &HostInfo{}
	host_data.hostId = id
	host_data.connectAddress = net.IP(GetServerAddress(id))
	host_data.port = int(port)
	host_data.peer = make([]int32, 4)
	for i := 1; i <= 3; i++ {
		if i != int(id) {
			host_data.peer[i] = int32(i)
		} else {
			host_data.peer[i] = int32(-1)
		}
	}
	hd.HostInfo[0] = host_data
	return host_data
}

func GetServerAddress(peerId int32) string {
	BASE_IP := "0.0.0.0"
	port := ":1200" + strconv.Itoa(int(peerId))
	addr := BASE_IP + port
	return addr
}

func (hd *InMemoryHostData) GetHostData() *HostInfo {
	host_data := hd.HostInfo[0]
	return host_data
}
func (hd *InMemoryHostData) GetHostId() int32 {
	host_id := hd.HostInfo[0].hostId
	return host_id
}
