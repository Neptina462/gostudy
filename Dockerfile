FROM golang

# 导入pingtest文件
COPY ./pingtest /webserve/pingtest
COPY ./pkg /webserve/pkg

# 导入tcp传输文件
COPY ./tcpclient /tcpfilesender/tcpclient
COPY ./tcpserve /tcpfilesender/tcpserve

# net-tools为安装网络工具
# inetutils-ping为ping工具
RUN apt-get update\
    && apt-get install net-tools\       
    && apt-get install inetutils-ping\
    && 
