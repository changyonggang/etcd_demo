package main

import (
	"context"
	"fmt"
	"github.com/changyonggang/etcd_demo/etcd_client"
	"github.com/changyonggang/etcd_demo/master"
	"github.com/changyonggang/etcd_demo/svc"
	"github.com/changyonggang/etcd_demo/utils"
	"github.com/changyonggang/etcd_demo/worker"
	"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"log"
	"sync"
)

/*
@Desc :
@Time : 2020/2/12 9:58 下午
@Author : Chang yg
@File : main
*/


func main() {
	// 1. init etcd client
	etcd_client.InitClient()
	defer etcd_client.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	// 2. 注册当前节点
	nodeId := utils.GetNodeId()
	fmt.Printf("this node id is %s \n", nodeId)
	go svc.RegisteSvc(etcd_client.EtcdClient, nodeId)

	// 3. go leader election
	go leaderElection(etcd_client.EtcdClient, nodeId)

	wg.Wait()
}

func leaderElection(cli *clientv3.Client, nodeId string) {
	s, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	elect := concurrency.NewElection(s, svc.SVC_ELECTION)
	electChannel := make(chan *concurrency.Election, 1)
	go func() {
		if err := elect.Campaign(context.Background(), nodeId); err != nil {
			log.Fatal(err)
		}
		electChannel <- elect
	}()

	// leader节点和slaver节点都要work
	go worker.SlaveWorker(cli, nodeId)

	// 竞选为leader  rpc  scheduler
	elect = <- electChannel
	master.MasterWorker(cli, nodeId)
}
