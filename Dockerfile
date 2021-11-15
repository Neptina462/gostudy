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

# 导入pingtest文件
COPY ./pingtest /webserve/pingtest

# 导入tcp传输文件，client为发送端，server为服务监听接收端
COPY ./tcpclient /tcpfilesender/tcpclient
COPY ./tcpserver /tcpfilesender/tcpserver

# 导入pkg文件到GOPATH目录
COPY ./gopath /go

# 导入GOROOT/src下的库
# GR为宿主机goroot中需要导入的新的库，需要手动更新
COPY ./GR /usr/local/go/src

# net-tools为安装网络工具
# inetutils-ping为ping工具
RUN apt-get update\
    && apt-get install net-tools\       
    && apt-get install inetutils-ping
