package etcd_client

import (
	"context"
	"github.com/changyonggang/etcd_demo/svc"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"testing"
)

/*
@Desc :
@Time : 2020/6/11 3:30 下午
@Author : Chang yg
@File : cli_test
*/

func TestInitClient(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{ Endpoints: []string{etcdUrl} })
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.Put(context.Background(), "name", "changyonggang")
	if err != nil {
		fmt.Printf("error %v\n", err)
	}
	fmt.Printf("resp  %v\n", resp.PrevKv)
	resp2 ,err := cli.Get(context.Background(), "name")
	if err != nil {
		fmt.Printf("error %v\n", err)
	}
	fmt.Printf("resp  %v\n", resp2.Kvs)
}

func TestInitClient2(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{ Endpoints: []string{etcdUrl} })
	if err != nil {
		log.Fatal(err)
	}
	kv := clientv3.NewKV(cli)
	resp2 ,err := kv.Get(context.Background(), svc.All_TASK_ROOT, clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("error %v\n", err)
	}
	fmt.Printf("resp  %v\n", resp2.Kvs)
}



