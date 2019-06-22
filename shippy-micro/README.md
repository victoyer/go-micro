# go-micro微服务
## 创建运行环境
```$xslt
1、mkdir tmp && cd tmp && go mod init tmp
2、go mod tidy
```
## 编译
### consignment-service:
```$xslt
cd consignment-service && make build && cd ..
```
### vessel-service:
```$xslt
cd vessel-service && make build && cd ..
```
### user-service:
```$xslt
cd user-service && make build && cd ..
```
### email-service:
```$xslt
cd email-service && make build && cd ..
```
### consignment-client:
```$xslt
cd consignment-client && make build && cd ..
```
### user-client:
```$xslt
cd user-client && make build && cd ..
```
## 执行 (在docker-compose.yaml文件所在目录)
```$xslt
1、docker stop $(docker ps -aq)
2、docker rm $(docker ps -aq)
3、docker rmi $(docker images -q)
4、docker-compose up --build
5、docker-compose run 微服务名称
...
```
ps: make 命令在mac下直接执行，如果是Windows需要[安装](http://gnuwin32.sourceforge.net/packages/make.htm)

