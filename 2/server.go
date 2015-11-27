package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
)

type Thing struct{}

func (t *Thing) ReverseWords(arg *Arg, reply *Reply) error {
	str := arg.Str
	parts := strings.Split(str, " ")
	out := ""
	for _, s := range parts {
		out = s + " " + out
	}
	reply.Str = out
	return nil
}

func main() {
	thing := new(Thing)
	rpc.Register(thing)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":54321")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)
}
