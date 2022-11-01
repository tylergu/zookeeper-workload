package zookeeperclient

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

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
		return fmt.Errorf("Failed to connect to zookeeper: %s, Reason: %v", zkUri, err)
	}
	client.conn = conn
	return nil
}

func (client *DefaultZookeeperClient) Close() {
	client.conn.Close()
}
