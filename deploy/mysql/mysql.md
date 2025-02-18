```sh
docker run -d \
  --name mysql-server \
  --restart=always \
  --privileged=true \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=Mygo*Mall6379 \
  -v /opt/mysql_data:/var/lib/mysql \
  -v /opt/mysql_config:/etc/mysql/conf.d \
  -v /opt/mysql_log:/var/log/mysql \
  -v /etc/localtime:/etc/localtime:ro
  mysql:8.0.33
```

| 选项                                     | 作用                                       |
| ---------------------------------------- | ------------------------------------------ |
| `-d`                                     | **后台运行** 容器                          |
| `--name mysql-server`                    | 设置容器名称为 `mysql-server`              |
| `--restart=always`                       | 容器异常退出时自动重启                     |
| `-p 3306:3306`                           | **映射端口**（建议限制访问）               |
| `-e MYSQL_ROOT_PASSWORD=StrongPass123!`  | **安全设置 root 密码**                     |
| `-e MYSQL_DATABASE=app_db`               | **初始化数据库 `app_db`**                  |
| `-e MYSQL_USER=app_user`                 | **创建业务数据库用户 `app_user`**          |
| `-e MYSQL_PASSWORD=AppUserPass123!`      | **为 `app_user` 设定密码**                 |
| `-v /opt/mysql_data:/var/lib/mysql`      | **数据持久化**，防止容器重启后数据丢失     |
| `-v /opt/mysql_config:/etc/mysql/conf.d` | **挂载配置文件**，方便管理                 |
| `--network my_network`                   | **使用 Docker 自定义网络**（防止外部访问） |

```sql
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'Mygo*Mall6379';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
```

