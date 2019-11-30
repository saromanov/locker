package zk

import (
	"fmt"
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/saromanov/locker/pkg/lock"
)

type zkLock struct {
	client *zk.Conn
}

func New(conf *lock.Config) lock.Locker {
	c, _, err := zk.Connect([]string{conf.Address}, time.Second)
	if err != nil {
		log.WithError(err).Error("unable connect to ZooKeeper")
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
		log.WithError(err).WithField("lock", "ZooKeeper").Error("unable connect to make lock")
		return false
	}

	return true
}

func (t *zkLock) Unlock() {
	l := zk.NewLock(t.client, "/lock", zk.WorldACL(zk.PermAll))
	l.Unlock()
}
