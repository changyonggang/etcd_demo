package etcd_client

import (
	"github.com/coreos/etcd/clientv3"
	"log"
)

/*
@Desc :
@Time : 2020/3/4 7:01 下午
@Author : Chang yg
@File : etcdCli
*/

const etcdUrl = ""

var EtcdClient  *clientv3.Client

func InitClient() {
	var err error
	EtcdClient, err = clientv3.New(clientv3.Config{ Endpoints: []string{etcdUrl} })
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	EtcdClient.Close()
}
