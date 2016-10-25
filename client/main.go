package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func main() {
	connection, err := net.Dial("unix", "/Users/patrickod/tmp/pilot.sock")
	if err != nil {
		log.Fatal("Unable to connect to socket: ", err)
	}

	client := rpc.NewClient(connection)

	var result *bool
	err = client.Call("Daemon.SyncClocks", 0, &result)
	if err != nil {
		log.Fatal("Error calling Daemon.SyncClocks: ", err)
	}

	fmt.Printf("Daemon.SyncClocks execution result: %v\n", *result)
}
