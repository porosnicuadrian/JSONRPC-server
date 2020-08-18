package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

var client *rpc.Client

func TestStartServer(t *testing.T) {
	go StartServer()
}

func TestConnClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	client = jsonrpc.NewClient(conn)
}

func TestWriteArgs(t *testing.T) {
	var reply string
	path := "numere.txt"
	args := &ArgsWrite{10, path}
	if err := client.Call("MyServer.WriteArgs", args, &reply); err != nil {
		t.Error(err)
	} else if reply != "TRUE" {
		t.Errorf("Expecting TRUE")
	}
}

func TestReadArgs(t *testing.T) {
	var reply int
	file := "numere.txt"
	args := &ArgsRead{file}
	if err := client.Call("MyServer.ReadArgs", args, &reply); err != nil {
		t.Error(err)
	} else if reply != 10 {
		t.Errorf("Expecting 10")
	}
	fmt.Println(reply)
}

func TestSum(t *testing.T) {
	var reply int
	args := &ArgsSum{6, 4}
	if err := client.Call("MyServer.Sum", args, &reply); err != nil {
		t.Error(err)
	} else if reply != 10 {
		t.Errorf("Expecting 10")
	}
	var replyWrite string
	filewrite := "suma.txt"
	sumwrite := &ArgsWrite{reply, filewrite}
	if err := client.Call("MyServer.WriteArgs", sumwrite, &replyWrite); err != nil {
		t.Error(err)
	} else if replyWrite != "TRUE" {
		t.Errorf("Expecting true")
	}
	var replyRead int
	sumread := &ArgsRead{filewrite}
	if err := client.Call("MyServer.ReadArgs", sumread, &replyRead); err != nil {
		t.Error(err)
	} else if replyRead != 10 {
		t.Errorf("Expecting 10")
	}
}
