package handler

import (
	"context"
	"fmt"
	"github.com/changyonggang/etcd_demo/etcd_client"
	"github.com/changyonggang/etcd_demo/svc"
)

/*
@Desc :
@Time : 2020/6/11 4:58 下午
@Author : Chang yg
@File : handler
*/


func AddTask(groupId, taskId string) error {
	_, err := etcd_client.EtcdClient.Put(context.Background(), svc.All_TASK_ROOT + groupId + "/" + taskId, fmt.Sprintf("%s-%s", groupId, taskId))
	return err
}

func DelTask(groupId, taskId string) error {
	_, err := etcd_client.EtcdClient.Delete(context.Background(), svc.All_TASK_ROOT + groupId + "/" + taskId)
	return err
}

