

# Iris





#### 3-10
xorm
```
/Users/lixiongjie/go/src/xorm.io/cmd/xorm/xorm reverse mysql "root:123456@tcp(127.0.0.1:3306)/cus?charset=utf8" templates/goxorm

[亲测GO环境搭建，理解go build、go install、go get](https://blog.csdn.net/zhangliangzi/article/details/77914943)

export GOPROXY=https://goproxy.cn
我以前用的是 https://goproxy.io 这个，但速度太慢而且还总动不动超时。。。。 (o_ _)ﾉ 。。。。最近在用 https://goproxy.cn 感觉很不错，基本没出现过什么问题，按作者说的是开发完之后会搬到国内来，到时候应该会更快吧

```

#### 3-9


go get github.com/go-xorm/cmd/xorm

```shell

cd /Users/lixiongjie/go/src/github.com/go-xorm/cmd/xorm

xorm reverse mysql "root:123456@tcp(127.0.0.1:3306)/?charset=utf8" templates/goxorm/


```




#### 3-8

mac 查看端口占用并
lsof -i :8080 | awk 'NR==2{print}' | awk '{print $2}'| xargs kill -9
lsof -i :8080
kill -9

[MacOS 上Golang Delve 调试填坑](https://www.jianshu.com/p/137854be2458)
[MySQL主从同步配置](https://blog.csdn.net/qq_41782425/article/details/88621138)
[Mac LiteIde 环境 代码提示](https://www.jianshu.com/p/46f425ea4fb6)
[GO LiteIDE 使用](https://studygolang.com/articles/5996)

#### 3-7

> MYSQL 主从配置

```yaml
docker run -p 3306:3306 -d -e MYSQL_ROOT_PASSWORD=123456 -v /Volumes/C/JAVA/mysqldb:/var/lib/mysql 64e
docker run -p 3307:3306 -d -e MYSQL_ROOT_PASSWORD=123456 -v /Volumes/C/JAVA/mysqldb2:/var/lib/mysql b3b

docker exec -it a04 bash



vim /etc/mysql/my.cnf


docker inspect -f '{{.Name}}-{{.NetworkSettings.IPAddress}}' $(docker ps -q)

1.在master中创建同步用户
use mysql;

select * from user;

grant REPLICATION slave,REPLICATION CLIENT on *.* to 'slave'@'%';
or
grant replication slave on *.* to 'slave'@'%'   ;


FLUSH PRIVILEGES;


2.在从库通过slave账号连接主库
mysql -h172.17.0.3 -P3306 -uslave -p123456

3.主库查询show master status;

4.从库配置
reset slave;

stop slave;

change master to master_host='172.17.0.3',master_user='slave',master_password='123456',master_port=3306
,master_log_file='mysql-bin.000002',master_log_pos=2206,master_connect_retry=30;

START SLAVE;
show slave status;


5.切记要保证没有数据库否则设置了也会同步失败

```


#### 3-6

无法安装iris是因为被墙了

[goproxy和go modules的初步使用](https://blog.csdn.net/qq_42403866/article/details/93654421)

 可以下只是很慢，缺少的包用git clone下载到src里


#### 3-5

> 简单队列 & 工作模式

工作模式有负载的作用，一个消息只能被一个消费者消费

> 订阅模式

多个消费者同时消费，队列绑定交换机，一个消息多个消费者

> 路由模式

生产端指定消费端进行消费

> topic 模式

[反过来选](https://www.cnblogs.com/LUA123/p/8477387.html)

> rabbitmq

docker pull rabbitmq:management

```yaml

docker run --name rabbitmq -d -p 15672:15672 -p 5672:5672 rabbitmq:management
docker stop rabbitmq
docker start rabbitmq
```

>控制台信息

启动容器后，可以浏览器中访问http://localhost:15672来查看控制台信息。

RabbitMQ默认的用户名：guest，密码：guest

>  go get github.com/streadway/amqp




```

```