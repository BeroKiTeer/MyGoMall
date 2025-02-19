# 三主三从

```sh
# 第一台
docker run -d --name redis-node-1 --net host --privileged=true -v /data/redis/share/redis-node-1:/data redis:7 --cluster-enabled yes --appendonly yes --port 6379
# 第二台
docker run -d --name redis-node-2 --net host --privileged=true -v /data/redis/share/redis-node-2:/data redis:7 --cluster-enabled yes --appendonly yes --port 6380
# 第三台
docker run -d --name redis-node-3 --net host --privileged=true -v /data/redis/share/redis-node-3:/data redis:7 --cluster-enabled yes --appendonly yes --port 6381
# 第四台
docker run -d --name redis-node-4 --net host --privileged=true -v /data/redis/share/redis-node-4:/data redis:7 --cluster-enabled yes --appendonly yes --port 6382
# 第五台
docker run -d --name redis-node-5 --net host --privileged=true -v /data/redis/share/redis-node-5:/data redis:7 --cluster-enabled yes --appendonly yes --port 6383
# 第六台
docker run -d --name redis-node-6 --net host --privileged=true -v /data/redis/share/redis-node-6:/data redis:7 --cluster-enabled yes --appendonly yes --port 6384

```

+ -d：表示以 “后台模式” 运行容器。
+ –net host：容器将与宿主机共享相同的网络接口和 IP 地址。
+ –privileged=true：启用容器的特权模式。
+ -v /data/redis/share/redis-node-1:/data：卷挂载，将宿主机的目录 /data/redis/share/redis-node-1 映射到容器内的 /data 目录。
+ –appendonly yes：启用 AOF (Append Only File) 持久化模式。

## 启动Redis集群

```sh
docker start redis-node-1 redis-node-2 redis-node-3 redis-node-4 redis-node-5 redis-node-6
```

将6个Redis容器搭建成三主三从集群模式

首先进入其中一个容器

```sh
docker exec -it redis-node-1 /bin/bash
```

搭建三主三从集群模式，一台主机对应一台从机

```shell
redis-cli --cluster create 172.31.0.2:6379 172.31.0.2:6380 172.31.0.2:6381 172.31.0.2:6382 172.31.0.2:6383 172.31.0.2:6384 --cluster-replicas 1
```

## 检查集群是否创建成功

启动Redis容器命令：

```sh
docker start redis-node-1 redis-node-2 redis-node-3 redis-node-4 redis-node-5 redis-node-6
```

进入其中一个Redis容器内部：

```shell
docker exec -it redis-node-1 /bin/bash
# 退出命令：exit
```

以集群的方式进入redis-node-1对应端口的客户端：

```sh
redis-cli -p 6379 -c
# 推出命令：quit
```

查看集群信息：

```bash
cluster info
```