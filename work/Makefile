export tag=v1.0

root:
    export ROOT=github.com/feiwang137/golang

build:
	echo "build httpServer image"
	GOOS=linux GOARCH=amd64 go build -o bin/httpServer ./main.go
	docker build -t wangfei1003/cncamp-httpserver:${tag} ./

push:
	echo "push image to docker hub."
	docker push wangfei1003/cncamp-httpserver:${tag}

run:
	echo "create httpServer container"
	docker run -d --name httpServer -p 80:80  wangfei1003/cncamp-httpserver:${tag}