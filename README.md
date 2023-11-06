# Api Gateway



## Getting started

```
1、编译
GOOS=linux  GOARCH=amd64 CGO_ENABLED=0 go build  -o recovery-tool  main.go

2、上传到主节点
rz -y 

3、启动命令
//开发环境
./recovery-tool -env=dev -c=config.yaml

//测试环境
./recovery-tool -env=test -c=config_test.yaml

//正式环境
./recovery-tool -env=online -c=config.encrypted

4、更新服务
ps -ef|grep gate
kill -1 923013

5、日志
/log

```
