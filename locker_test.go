package dlock

import (
	"testing"

	"github.com/wl955/log"
)

func init() {
	Init([]string{"http://127.0.0.1:2379"})
}

func TestNewLocker(t *testing.T) {
	locker, e := NewLocker("/lock", 1)
	if e != nil {
		t.Fatal(e)
	}
	locker.Lock()
	defer locker.Unlock()
	t.Log("acquired lock for s1")
}

func TestNewLocker2(t *testing.T) {
	m1, e := NewLocker("/my-lock/", 60)
	if e != nil {
		t.Fatal(e)
	}

	m2, e := NewLocker("/my-lock/", 60)
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

func TestLockVer(t *testing.T) {
	test()
}

func Test1(t *testing.T) {
	log.Info("fafa\n")
	log.Info("fafa\n")
}
