package main

import (
	"fmt"
	"time"

	zkcli "github.com/samuel/go-zookeeper/zk"
)

func main() {
	//c, _, err := zk.Connect([]string{"14.22.10.93"}, time.Second) //*10)
	c, _, err := zkcli.Connect([]string{"14.29.105.178"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	authData := fmt.Sprintf("%s:%s", "admin", "admin")
	err = c.AddAuth("sasl", []byte(authData))
	children, stat, ch, err := c.ChildrenW("/zookeeper")
	if err != nil {
		panic(err)
	}
	fmt.Println(children)
	fmt.Printf("%+v %+v\n", children, stat)
	e := <-ch
	fmt.Printf("%+v\n", e)
}
