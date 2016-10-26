package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
)

type Daemon struct{}

func (d *Daemon) SyncClocks(time int, result *bool) error {
	fmt.Println("Synced the clocks")
	*result = true
	return nil
}

func main() {
	path := "/tmp/pilot.sock"
	daemon := new(Daemon)

	rpc.Register(daemon)
	listener, err := net.Listen("unix", path)
	if err != nil {
		log.Fatal("Unable to listen ", err)
	}

	// Make sure that we tear down the socket appropriately
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		listener.Close()
		os.Remove(path)
	}()

	rpc.Accept(listener)
}
