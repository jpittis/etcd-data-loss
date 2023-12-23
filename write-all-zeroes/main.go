package main

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	_, err = client.Put(context.Background(), "all-zeroes", string(buf))
	if err != nil {
		log.Fatal(err)
	}
}
