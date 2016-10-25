package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Daemon struct{}

func (d *Daemon) SyncClocks(time int, result *bool) error {
	fmt.Println("Synced the clocks")
	*result = true
	return nil
}

func main() {

	daemon := new(Daemon)
	rpc.Register(daemon)

	listener, err := net.Listen("unix", "/Users/patrickod/tmp/pilot.sock")
	if err != nil {
		log.Fatal("Unable to listen on /Users/patrickod/tmp/pilot.sock: ", err)
	}

	rpc.Accept(listener)
}
