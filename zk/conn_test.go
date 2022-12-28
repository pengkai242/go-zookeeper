package zk

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestRecurringReAuthHang(t *testing.T) {
	t.Skip("Race condition in test")

	sessionTimeout := 2 * time.Second

	finish := make(chan struct{})
	defer close(finish)
	go func() {
		select {
		case <-finish:
			return
		case <-time.After(5 * sessionTimeout):
			panic("expected not hang")
		}
	}()

	zkC, err := StartTestCluster(t, 2, ioutil.Discard, ioutil.Discard)
	if err != nil {
		panic(err)
	}
	defer zkC.Stop()

	conn, evtC, err := zkC.ConnectAll()
	if err != nil {
		panic(err)
	}
	for conn.State() != StateHasSession {
		time.Sleep(50 * time.Millisecond)
	}

	go func() {
		for range evtC {
		}
	}()

	// Add auth.
	conn.AddAuth("digest", []byte("test:test"))

	currentServer := conn.Server()
	conn.debugCloseRecvLoop = true
	//conn.debugReauthDone = make(chan struct{})
	zkC.StopServer(currentServer)
	// wait connect to new zookeeper.
	for conn.Server() == currentServer && conn.State() != StateHasSession {
		time.Sleep(100 * time.Millisecond)
	}

	//<-conn.debugReauthDone
}

func TestSasl2(t *testing.T) {
	connectString := "14.22.10.93:12"
	t.Log(connectString[:])
	c, _, err := Connect([]string{"14.22.10.93"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	children, stat, err := c.Children("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)
}
