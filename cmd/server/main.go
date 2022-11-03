package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tylergu/zookeeper-workload/zookeeperclient"
)

func main() {
	errC := make(chan error)
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		errC <- startServer()
	}()

	select {
	case err := <-errC:
		fmt.Printf("Error: %v", err)
	case sig := <-sigC:
		fmt.Printf("Signal: %v", sig)
	}

	os.Exit(0)
}

func startServer() error {
	for {
		client := zookeeperclient.DefaultZookeeperClient{}
		err := client.Connect(zookeeperclient.GetZkServiceUri())
		if err != nil {
			fmt.Printf("Failed to connect to zookeeper: %v", err)
			continue
		}

		err = client.Create("/test", []byte("test"))
		if err != nil {
			return err
		} else {
			break
		}
	}

	for {
		fmt.Println("ALIVE")
		time.Sleep(5 * time.Second)
	}
}
