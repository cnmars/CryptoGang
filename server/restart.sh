#!/bin/bash
echo "go build"
# go mod tidy
go build -o golddigger
# go build -o golddigger main.go
echo "kill golddigger service"
killall golddigger # kill golddigger service
# nohup ./golddigger start -c=config/config.yaml >> access.log 2>&1 & #后台启动服务将日志写入access.log文件
./golddigger start
echo "run golddigger success" 
ps -aux | grep golddigger