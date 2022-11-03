package main

import (
	"fmt"

	"github.com/tylergu/zookeeper-workload/zookeeperclient"
)

func main() {
	client := zookeeperclient.DefaultZookeeperClient{}
	err := client.Connect(zookeeperclient.GetZkServiceUri())
	if err != nil {
		panic(err)
	}

	data, err := client.Get("/test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", data)
}
