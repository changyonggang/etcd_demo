package master

import (
	"context"
	"fmt"
	"github.com/changyonggang/etcd_demo/http_svc"
	"github.com/changyonggang/etcd_demo/schedule"
	"github.com/changyonggang/etcd_demo/svc"
	"github.com/coreos/etcd/clientv3"
	"net/http"
	"time"
)

/*
@Desc :
@Time : 2020/3/4 3:00 下午
@Author : Chang yg
@File : master
*/

var AliveNode map[string]string

func init() {
	AliveNode = make(map[string]string)
}

func MasterWorker(cli * clientv3.Client, nodeId string) {
	fmt.Printf("master %s worker start ... ... \n", nodeId)

	go printAliveNode()
	// 1. init rpc http server
	go initHttpSvc(nodeId)

	// 2. service find
	go nodeFind(cli)

	// 3. 处理上一届master留下的task列表

	// 4. load balance
	go loadBalance(cli)

	// 5. task schdule
	schedule.Schedule(cli)
}

func initHttpSvc(port string) {
	http.HandleFunc("/ping", http_svc.PingHandler)
	http.HandleFunc("/add", http_svc.AddHandler)
	http.HandleFunc("/delete", http_svc.DelHandler)
	_ = http.ListenAndServe("0.0.0.0:8000", nil)
}

func printAliveNode() {
	for {
		for k, _ := range AliveNode  {
			fmt.Printf("alive node %s \n", k)
		}
		time.Sleep(30 * time.Second)
	}
}

func loadBalance(cli *clientv3.Client) {
	// TODO::
}

func nodeFind(cli *clientv3.Client) {
	rch := cli.Watch(context.Background(), svc.SVC_NODE_ROOT, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.IsModify() {
				//fmt.Printf("is modify \n")
				updateAliveNode(fmt.Sprintf("%s", ev.Kv.Key))
			}
			if ev.IsCreate() {
				fmt.Printf("node %s add cluster ...\n", ev.Kv.Key)
				updateAliveNode(fmt.Sprintf("%s", ev.Kv.Key))
			}
			if ev.Type == 1 {
				fmt.Printf("node %s dead ..\n", ev.Kv.Key)
				DeleteAliveNode(fmt.Sprintf("%s", ev.Kv.Key))
			}
		}
	}
}

func DeleteAliveNode(node string) {
	delete(AliveNode, node)
}

func updateAliveNode(node string) {
	if _, ok := AliveNode[node]; !ok {
		AliveNode[node] = time.Now().String()
	}
}