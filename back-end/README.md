# iBooking



#### 部署至docker

```
docker build -t ibooking .
# 后台运行
docker run -d -p 8800:8800 ibooking
# 前台运行
docker run -p 8800:8800 ibooking
```

#### 从docker删除

```
docker stop psid
docker rm psid
docker rmi ibooking
```
