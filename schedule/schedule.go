package schedule

import (
	"context"
	"github.com/changyonggang/etcd_demo/svc"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"strings"
)

/*
@Desc :
@Time : 2020/3/4 6:42 下午
@Author : Chang yg
@File : schedule
*/


func Schedule(cli *clientv3.Client) {
	rch := cli.Watch(context.Background(), svc.All_TASK_ROOT, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.IsModify() {
				// updateTaskMap(fmt.Sprintf("%s", ev.Kv.Value))
			}
			if ev.IsCreate() {
				fmt.Printf("task %s add cluster ...\n", ev.Kv.Key)
				assignTask(cli, string(ev.Kv.Key), string(ev.Kv.Value))
			}
			// 删除
			if ev.Type == 1 {
				fmt.Printf("delete %s task ..\n", ev.Kv.Key)
				delTask(cli, string(ev.Kv.Key))
			}
		}
	}
}

func assignTask(cli *clientv3.Client, taskNode, taskValue string) {
	fmt.Println(svc.TASK_NODE_ROOT + strings.TrimLeft(taskNode, svc.All_TASK_ROOT))
	_, err := cli.Put(context.Background(), svc.TASK_NODE_ROOT + strings.TrimLeft(taskNode, svc.All_TASK_ROOT), taskValue)
	if err != nil {
		fmt.Println(err)
	}
}


func delTask(cli *clientv3.Client, taskNode string) {
	fmt.Println(svc.TASK_NODE_ROOT + strings.TrimLeft(taskNode, svc.All_TASK_ROOT))
	_, err := cli.Delete(context.Background(), svc.TASK_NODE_ROOT + strings.TrimLeft(taskNode, svc.All_TASK_ROOT))
	if err != nil {
		fmt.Println(err)
	}
}