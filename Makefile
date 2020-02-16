all: local

local:
	GOOS=linux GOARCH=amd64 go build  -o=latte .

build:
	sudo docker build --no-cache . -t registry.cn-hangzhou.aliyuncs.com/njput-isl/latte

push:
	sudo docker push registry.cn-hangzhou.aliyuncs.com/njupt-isl/latte

format:
	sudo gofmt -l -w .
clean:
	sudo rm -f scv