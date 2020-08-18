package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strconv"
)

type MyServer struct{}

func (srv *MyServer) Sum(args ArgsSum, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func (srv *MyServer) WriteArgs(args ArgsWrite, reply *string) error {
	file, err := os.Create(args.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(file, "%d", args.X)
	*reply = "TRUE"
	return nil
}

func (srv *MyServer) ReadArgs(args ArgsRead, reply *int) error {
	file, err := ioutil.ReadFile(args.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	number, err := strconv.Atoi(string(file))
	*reply = number
	return nil
}

func StartServer() {
	fmt.Println("Start server")
	srv := new(MyServer)
	server := rpc.NewServer()
	server.Register(srv)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := listener.Accept()
		fmt.Println("Connection established")
		if err != nil {
			log.Fatal("Conecction error")
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(connection))
		defer connection.Close()
		fmt.Println("Server close")
	}
}
