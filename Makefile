gen:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:.

server_main: Server/server_main.go
	go build $<

client_main: main.go
	go build $<
client:
	./main
server: server1 server2 server3
server1:
	./server_main -id 1
server2:
	./server_main -id 2
server3:
	./server_main -id 3
clean:
	rm go.mod
	rm go.sum