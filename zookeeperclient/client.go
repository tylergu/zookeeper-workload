package zookeeperclient

import (
	"fmt"
	"strings"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var ZK_CLUSTER_NAME string = "test-cluster"

type ZookeeperClient interface {
	Connect(string) error
	Close()
}

type DefaultZookeeperClient struct {
	conn *zk.Conn
}

func (client *DefaultZookeeperClient) Connect(zkUri string) (err error) {
	host := []string{zkUri}
	conn, _, err := zk.Connect(host, time.Second*5)
	if err != nil {
		return fmt.Errorf("failed to connect to zookeeper: %s, Reason: %v", zkUri, err)
	}
	client.conn = conn
	return nil
}

func (client *DefaultZookeeperClient) Create(path string, data []byte) error {
	paths := strings.Split(path, "/")
	pathLength := len(paths)
	var parentPath string
	for i := 1; i < pathLength-1; i++ {
		parentPath += "/" + paths[i]
		if _, err := client.conn.Create(parentPath, nil, 0, zk.WorldACL(zk.PermAll)); err != nil {
			return fmt.Errorf("error creating parent zkNode: %s: %v", parentPath, err)
		}
	}

	childNode := parentPath + "/" + paths[pathLength-1]
	if _, err := client.conn.Create(childNode, []byte(data), 0, zk.WorldACL(zk.PermAll)); err != nil {
		return fmt.Errorf("error creating sub zkNode: %s: %v", childNode, err)
	}
	return nil
}

func (client *DefaultZookeeperClient) Get(path string) ([]byte, error) {
	data, _, err := client.conn.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error getting zkNode: %s: %v", path, err)
	}
	return data, nil
}

func (client *DefaultZookeeperClient) Close() {
	client.conn.Close()
}

func GetZkServiceUri() string {
	return fmt.Sprintf("%s-client", ZK_CLUSTER_NAME) + ":tcp-client"
}
