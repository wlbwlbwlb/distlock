package distlock

import (
	"fmt"
	"os"
	"testing"

	"go.etcd.io/etcd/client/v3/concurrency"
)

func init() {
	Init(WithEndpoints([]string{"http://127.0.0.1:2379"}))
}

func TestNewLocker(t *testing.T) {
	m1, e := New("/my-lock/", 60)
	if e != nil {
		t.Fatal(e)
	}

	m2, e := New("/my-lock/", 60)
	if e != nil {
		t.Fatal(e)
	}

	m1.Lock()
	t.Log("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// wait until s1 is locks /my-lock/
		m2.Lock()
	}()

	m1.Unlock()
	t.Log("released lock for s1")

	<-m2Locked
	t.Log("acquired lock for s2")
}

const (
	// These const values might be need adjustment.
	nrGarbageObjects = 100 * 1000 * 1000
	sessionTTL       = 1
)

func TestLockVer(t *testing.T) {
	session, err := concurrency.NewSession(client, concurrency.WithTTL(sessionTTL))
	if err != nil {
		fmt.Printf("failed to create a session: %s\n", err)
		os.Exit(1)
	}

	locker := concurrency.NewLocker(session, "/lock")
	locker.Lock()
	defer locker.Unlock()

	version := session.Lease()
	fmt.Printf("acquired lock, version: %d\n", version)
}
