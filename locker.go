package distlock

import (
	"errors"
	"sync"

	"github.com/wlbwlbwlb/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var _opts clientv3.Config

var client *clientv3.Client

func Init(opts ...Option) {
	custom := Options{}

	for _, o := range opts {
		o.apply(&custom)
	}

	var err error

	client, err = clientv3.New(
		_opts,
	)
	if err != nil {
		log.Fatalf("failed to create an etcd client: %s\n", err)
	}

	log.Info("creted etcd client\n")
}

func New(pfx string, ttl int) (locker sync.Locker, err error) {
	if nil == client {
		return nil, errors.New("init first")
	}
	// WithTTL configures the session's TTL in seconds.
	session, err := concurrency.NewSession(client, concurrency.WithTTL(ttl))
	if err != nil {
		return
	}

	locker = concurrency.NewLocker(session, pfx)

	return
}
