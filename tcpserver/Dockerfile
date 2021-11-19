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

# 导入tcpserver服务监听接收端
COPY . /tcpserver

# net-tools为安装网络工具
# inetutils-ping为ping工具
# download为挂载的接收区域
RUN apt-get update\
    && apt-get install net-tools\       
    && apt-get install inetutils-ping\
    && mkdir /download

#启动服务，默认8080端口
ENTRYPOINT ["go","run", "/tcpserver/main.go"]

# 启动docker时可以加入动态参数修改监听端口
# eg: docker run [container's name] -p :7070