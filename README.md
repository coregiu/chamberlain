[![Language](https://img.shields.io/badge/language-go-blue.svg)](README.md)
[![Language](https://img.shields.io/badge/language-vue-blue.svg)](README.md)
[![License](https://img.shields.io/badge/license-Apache2.0-brightgreen.svg)](LICENSE)

# Self website for code example and experies summary.

## Function:
1. Self blogs, synchronize from github repository and convert from markdown to html automaticly.
2. Self memorandum and diary, edit, display and search online.
3. Selt inputs management, analysis types of income and the trend of income.
4. Has completable system management, like user management, auth check, log management, system backup and resoter and so on.
5. At the last, I want to implements a game of my hometown named the cannon bombards the dogfaces.

## Get started:
1. Install golang, node, npm, git.
2. Clone code.
3. Compile the chamberlain_website and make docker image:
```shell
npm run build
docker build -t regiu:1.0 .
```
4. Compile the chamberlain_mtmt and make docker image:
```shell
docker build -t chamberlain:1.0 .
```
5. Set the image version to docker-compose.yml
6. Pull mysql image

7. Run docker compose:
```shell
# 一 使用docker stack启动集群
# 创建集群
docker swarm init --advertise-addr xx.xx.xx.xx

# 启动服务
docker stack deploy chamberlain --compose-file docker-compose.yaml
# 删除服务
docker stack remove chamberlain
# 查看服务
docker service ls
docker service logs vjdew7q0hmvr

# 导入数据
## 进入容器
docker exec -it 8e3537249b94 /bin/sh 
## 创建新库
mysql -u root -p
CREATE DATABASE chamberlain;

## 导入库表
mysql -u root -D chamberlain -p

# 二 单独启动
docker run -d --name redis -p 6379:6379 -e MYSQL_ROOT_PASSWORD=xxxxxx redis:latest
docker run -d --name database -v /home/coregiu/workspace/runtime/chamberlain/database:/var/lib/mysql -v /home/coregiu/workspace/runtime/chamberlain/database:/etc/mysql/conf.d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=xxxxxx mysql:latest
docker run -d --name chamberlain -v /var:/var -v /home/coregiu/workspace/runtime/chamberlain/chamberlain.yml:/opt/chamberlain/chamberlain.yml --link database:database --link redis:redis -p 8080:8080 -e MYSQL_ROOT_PASSWORD=xxxxxx -e REDIS_ROOT_PASSWORD=xxxxxx chamberlain:1.0
docker run -d --name regiu --link chamberlain:chamberlain -p 80:80 regiu:1.0
```
8. Docker exec the mysql container, init database and tables. The sql is chamberlain.sql.
9. Then you can open the page in browse and operate it.