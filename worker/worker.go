package worker

import (
	"context"
	"fmt"
	"github.com/changyonggang/etcd_demo/svc"
	"github.com/coreos/etcd/clientv3"
	"time"
)

/*
@Desc :
@Time : 2020/3/4 3:01 下午
@Author : Chang yg
@File : worker
*/
var taskMap map[string]string

func init() {
	taskMap = make(map[string]string)
}

func SlaveWorker(cli * clientv3.Client, nodeId string) {
	fmt.Printf("slave %s worker start ... ... \n", nodeId)
	go printRunningTask()
	taskFind(cli, nodeId)
}

func printRunningTask() {
	for {
		for k, v := range taskMap  {
			fmt.Printf("running task %s,\t %s \n", k, v)
		}
		if len(taskMap) > 0 {
			fmt.Println("running tasks end")
		} else {
			fmt.Println("no running task")
		}
		time.Sleep(20 * time.Second)
	}
}

func taskFind(cli *clientv3.Client, nodeId string) {
	rch := cli.Watch(context.Background(), svc.TASK_NODE_ROOT + nodeId, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.IsModify() {
				updateTaskMap(string(ev.Kv.Key), string(ev.Kv.Value))
			}
			if ev.IsCreate() {
				updateTaskMap(string(ev.Kv.Key), string(ev.Kv.Value))
			}
			if ev.Type == 1 {
				deleteTaskMap(string(ev.Kv.Key))
			}
		}
	}
}

func updateTaskMap(taskKey, taskValue string) {
	fmt.Println("update :" + taskKey)
	taskMap[taskKey] = taskValue
}


func deleteTaskMap(taskKey string) {
	fmt.Println("delete :" + taskKey)
	delete(taskMap, taskKey)
}

