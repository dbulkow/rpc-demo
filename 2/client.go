package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:54321")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Arg{Str: "some string that wants to be backwards"}
	var reply Reply

	err = client.Call("Thing.ReverseWords", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(reply.Str)
}
