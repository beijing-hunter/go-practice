package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)
func main()  {

	var(
		config clientv3.Config
	)

	config=clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5*time.Second,
	}

	client,err:=clientv3.New(config)

	if err!=nil{
		fmt.Println(err)
		return
	}

	client.Close()
}
