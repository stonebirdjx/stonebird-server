# 系统运行环境

## 创建docker网络
```bash
docker network create -d bridge stonebird-network
docker network ls
```

## mysql

### 安装mysql
```bash
docker pull mysql

mkdir -p /stonebird/data/mysql

docker run -d \
--restart=always \
--name stonebird-mysql \
--network stonebird-network \
-p 3306:3306 \
-v /stonebird/data/mysql:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=123456 \
mysql:latest
```

## 设置环境变量
```bash
	export STB_MYSQL_HOST=ip
	export STB_MYSQL_PORT=port
	export STB_MYSQL_USER=user
	export STB_MYSQL_PASSWD=passwd
```

## DDL
### 创建数据库
```bash
create database stonebird;
show databases;
```
