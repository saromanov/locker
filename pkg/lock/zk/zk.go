package zk

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"github.com/saromanov/locker/pkg/lock"
)

type zkLock struct {
	client *zk.Conn
}

func New() lock.Locker {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)
	if err != nil {
		fmt.Println("unable to connect to ZooKeeper")
		return nil
	}
	return &zkLock{
		client: c,
	}
}

func (t *zkLock) Lock() bool {
	l := zk.NewLock(t.client, "/lock", zk.WorldACL(zk.PermAll))
	err := l.Lock()
	if err != nil {
		fmt.Println("unable to make Lock: %v", err)
		return false
	}

	return true
}

func (t *zkLock) Unlock() {
	l := zk.NewLock(t.client, "/lock", zk.WorldACL(zk.PermAll))
	l.Unlock()
}
