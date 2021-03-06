package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

//发送文件到服务端
func SendFile(filePath string, fileSize int64, conn net.Conn) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	var count int64
	for {
		buf := make([]byte, 2048)
		//读取文件内容
		n, err := f.Read(buf)
		if err != nil && io.EOF == err {
			fmt.Println("文件传输完成")
			//告诉服务端结束文件接收
			conn.Write([]byte("finish"))
			return
		}
		//发送给服务端
		conn.Write(buf[:n])

		count += int64(n)
		sendPercent := float64(count) / float64(fileSize) * 100
		value := fmt.Sprintf("%.2f", sendPercent)
		//打印上传进度
		fmt.Printf("文件上传进度：%s/100.00\r", value)
	}
}

func main() {

	var server_ip string //服务器地址，ip+端口号
	var file_url string  //文件路径

	flag.StringVar(&server_ip, "p", "127.0.0.1:8080", "server's ip + port")
	flag.StringVar(&file_url, "f", "", "file's location")
	flag.Parse()

	//获取文件信息
	fileInfo, err := os.Stat(file_url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建客户端连接
	conn, err := net.Dial("tcp", server_ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//文件名称
	fileName := fileInfo.Name()
	//文件大小
	fileSize := fileInfo.Size()
	//发送文件名称到服务端
	conn.Write([]byte(fileName))
	buf := make([]byte, 2048)
	//读取服务端内容
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	revData := string(buf[:n])
	if revData == "ok" {
		//发送文件数据
		SendFile(file_url, fileSize, conn)
	}
}
