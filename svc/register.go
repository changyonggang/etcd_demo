package svc

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

const SVC_NODE_ROOT = "/etcd_demo/svc_node/"
const All_TASK_ROOT = "/etcd_demo/all_task/"
const TASK_NODE_ROOT = "/etcd_demo/task_node/"
const SVC_ELECTION = "/etcd_demo/election/"
/*
@Desc :
@Time : 2020/3/4 3:22 下午
@Author : Chang yg
@File : register
*/

func RegisteSvc(cli * clientv3.Client, nodeId string) {
	heartBeat(cli, nodeId)
}

func heartBeat(cli * clientv3.Client, nodeId string) {
	key := SVC_NODE_ROOT + nodeId // 用 pid 来标识每一个服务， 通常应该用 IP 等来标识。
	for {
		// etcd 之所以适合用来做服务发现，是因为它是带目录结构的。 注册一类服务，
		resp, _ := cli.Grant(context.TODO(), 20)
		ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
		_, err := cli.Put(ctx, key, time.Now().String(), clientv3.WithLease(resp.ID))
		if err != nil {
			log.Println("Error update workerInfo:", err)
		}
		// fmt.Printf("node:%s heart beat ...\n", nodeId)
		time.Sleep(time.Second * 10)
	}
}