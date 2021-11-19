# 基于golang官方包
# 部分参数：
# GOROOT=/usr/local/go
# GOPROXY="https://proxy.golang.org,direct"
# GOENV="/root/.config/go/env"（该目录需要自行建立）
# GOMODCACHE="/go/pkg/mod"
# GOPATH="/go"（pkg在此目录之下）
FROM golang

# DOCKER的GOPATH
ENV DOCKER_GOPATH /go  
# 宿主机的GOPATH
ENV LOCAL_GOPATH ./gopath

# 导入tcp传输文件，client为发送端，server为服务监听接收端
COPY . /tcpclient

# net-tools为安装网络工具
# inetutils-ping为ping工具
# /sender为挂载的发送区
RUN apt-get update\
    && apt-get install net-tools\       
    && apt-get install inetutils-ping\
    && mkdir /sender

# 启动服务，默认地址为：127.0.0.1:8080，默认文件为空
ENTRYPOINT ["go","run", "/tcpclient/main.go"]

# 启动docker时可以加入动态参数修改接收的服务器端地址，以及文件位置
# eg:docker start [container's name] -p 10.176.43.180:5050 -f ./arona.png