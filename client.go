package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	var reply int
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	args := &ArgsSum{2, 3}
	client := jsonrpc.NewClient(conn)
	client.Call("MyServer.Sum", args, &reply)
	fmt.Println(reply)

}
