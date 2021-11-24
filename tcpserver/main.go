package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
)

func Handler(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	//读取客户端发送的内容
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := string(buf[:n])
	//获取客户端ip+port
	addr := conn.RemoteAddr().String()
	fmt.Println(addr + ": 客户端传输的文件名为--" + fileName)
	//告诉客户端已经接收到文件名
	conn.Write([]byte("ok"))

	//保存的文件夹
	dirname := "/download/"
	//判断文件夹是否存在
	_, err0 := os.Stat(dirname)
	if err0 != nil {
		//若错误为不存在
		if os.IsNotExist(err0) {
			fmt.Println("文件夹不存在，即将创建")
			err1 := os.Mkdir(dirname, os.ModePerm)
			if err1 != nil {
				fmt.Println("文件夹创建失败！即将退出")
				fmt.Println(err1)
				return
			}
			fmt.Println("创建成功！")
		} else {
			fmt.Println(err0)
			return
		}
	}

	//创建文件
	f, err := os.Create(dirname + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	//循环接收客户端传递的文件内容
	for {
		buf := make([]byte, 2048)
		n, _ := conn.Read(buf)
		//结束协程
		if string(buf[:n]) == "finish" || n == 0 {
			fmt.Println(addr + ": 协程结束")
			runtime.Goexit()
		}
		f.Write(buf[:n])
	}

}

func main() {
	/* fmt.Println("输入需要监听的端口")
	var port_name string
	fmt.Scan(&port_name) */
	//创建tcp监听

	var port string

	flag.StringVar(&port, "p", ":8080", "port to be listened")
	flag.Parse()

	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	for {
		//阻塞等待客户端
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		//创建协程
		go Handler(conn)
	}
}
