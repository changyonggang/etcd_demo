

# etcd-demo介绍

## 文件树

```
.
├── ReadMe.md
├── etcd_client
│   ├── cli_test.go
│   └── etcdCli.go
├── go.mod
├── go.sum
├── handler
│   └── handler.go
├── http_svc
│   └── http_svc.go
├── main.go
├── master
│   └── master.go
├── schedule
│   └── schedule.go
├── svc
│   ├── register.go
│   └── service.go
├── utils
│   ├── utils.go
│   └── utils_test.go
└── worker
    └── worker.go

8 directories, 16 files
```

## 简介

etcd-demo根据etcd作为分布式协调器，完成了一个简单的分布式任务分发系统。master节点接收任务，并把任务分发给其他节点，为了节省节点，master节点除了完成任务调度的职责外，还同时作为一个worker存在。本demo根据进程id作为node的id来标识不同的节点，所以支持同一个机器模拟。

### done

1. leader选举
2. node 发现
3. master作为http server 接收任务和删除任务

### run

1. 编译

   ```shell
   go build -a -o demo
   ```

2. 在不同的shell窗口启动服务

   ```
   ./demo
   ```

3. 模拟不同的节点宕机验证leader选举功能

4. 验证任务分发功能

   1. 添加任务

      ```
      http://master:8000/add?groupId=<node-id>&&taskId=<task-id>
      ```

   2. 验证对应节点是否有对应的任务运行，即任务下发时候成功。

   3. 删除已有任务

      ```
      http://master:8000/delete?groupId=<node-id>&&taskId=<task-id>
      ```

   4. 验证对应节点任务是否下架。