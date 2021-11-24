# 基于golang官方包
# 部分参数：
# GOROOT=/usr/local/go
# GOPROXY="https://proxy.golang.org,direct"
# GOENV="/root/.config/go/env"（该目录需要自行建立）
# GOMODCACHE="/go/pkg/mod"
# GOPATH="/go"（pkg在此目录之下）
FROM golang

# 导入tcp传输文件，client为发送端，server为服务监听接收端
COPY . /serve

# /sender为挂载的发送区
RUN  mkdir /sender

# 启动服务，默认地址为：127.0.0.1:8080，默认文件为空
ENTRYPOINT ["go","run", "/serve/main.go" ]

# 启动docker时可以加入动态参数修改接收的服务器端地址，以及文件名称
# eg:docker run [image's name] -p 10.176.43.180:5050 -f arona.png