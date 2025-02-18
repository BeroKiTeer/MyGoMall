### 安装Prometheus

```bash
[root@Docker ~]# docker pull prom/prometheus
[root@Docker ~]# docker run -itd --name=prometheus --restart=always -p 9090:9090 prom/prometheus
```

### 安装Grafana

Grafana是一个跨平台开源的度量分析和可视化工具，可以通过将采集的数据查询然后可视化的展示，并及时通知。

创建挂载数据目录

```bash
mkdir /opt/grafana-storage
```

设置权限

```bash
chmod 777 -R /opt/grafana-storage
```

```bash
[root@Docker ~]# docker pull grafana/grafana
[root@Docker ~]# docker run -d \
-p 3000:3000 \
--restart=always \
--name=grafana \
-v /opt/grafana-storage:/var/lib/grafana \
grafana/grafana

```

### 安装Node_exporter

因为Prometheus 本身不具备监控功能，所以想要通过Prometheus 收集数据的话，需要安装对应的exporter。

```bash
[root@Docker ~]# docker pull prom/node-exporter
[root@Docker ~]# docker run -itd --name=node-exporter \
--restart=always \
-p 9100:9100 \
-v "/proc:/host/proc:ro" \
-v "/sys:/host/sys:ro" \
-v "/:/rootfs:ro" \
prom/node-exporter
```

