# 模块三作业

## 将 httpServer 应用容器化

```
git clone git@github.com:feiwang137/golang.git
cd golang/work

# 构建镜像
make build

# 将镜像推送到 docker 远程仓库
make push
```
## 本地启动 httpServer 

```
# 方式1： 通过Makefile 启动容器
make run

# 方式2： docker 直接启动
docker run -d --name httpServer -p 80:80  wangfei1003/cncamp-httpserver:v1.0
```

## 查看 httpServer 容器ip
```
pid=`lsns -t net|grep httpServer|awk '{print $4}'`
nsenter -t $pid -n ip addr
```