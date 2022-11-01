package main

import (
	"fmt"

	"github.com/tylergu/zookeeper-workload/zookeeperclient"
)

var ZK_CLUSTER_NAME string = "test-cluster"

func main() {
	client := zookeeperclient.DefaultZookeeperClient{}
	err := client.Connect(GetZkServiceUri())
	if err != nil {
		panic(err)
	}
}

func GetZkServiceUri() string {
	return fmt.Sprintf("%s-client", ZK_CLUSTER_NAME) + ":2181"
}
