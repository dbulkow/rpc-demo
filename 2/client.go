package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func ReverseWords(str string) string {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:54321")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Arg{Str: str}
	var reply Reply

	err = client.Call("Thing.ReverseWords", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	return reply.Str
}

func main() {
	val := ReverseWords("some string that wants to be backwards")
	fmt.Println(val)
}
