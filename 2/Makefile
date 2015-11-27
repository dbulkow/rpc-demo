all: client server

client: client.go common.go
	go build -v $^

server: server.go common.go
	go build -v $^
