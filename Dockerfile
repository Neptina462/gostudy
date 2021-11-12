FROM golang

# 导入pingtest文件
COPY ./pingtest /webserve/pingtest
COPY ./pkg /webserve/pkg

# 导入tcp传输文件，client为发送端，server为服务监听接收端
COPY ./tcpclient /tcpfilesender/tcpclient
COPY ./tcpserver /tcpfilesender/tcpserver

# net-tools为安装网络工具
# inetutils-ping为ping工具
RUN apt-get update\
    && apt-get install net-tools\       
    && apt-get install inetutils-ping
