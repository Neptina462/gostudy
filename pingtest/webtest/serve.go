package webtest

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strings"
)

func SayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数
	fmt.Println(r.Form) //服务器端输出信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello astaxie!")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, nil))
	} else if r.Method == "POST" {
		err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}

		//获取服务端ip
		serve_url, err_2 := GetOutBoundIP()
		if err_2 != nil {
			// handle error http.Error() for example
			log.Fatal("GetIPerr: ", err)
		}

		result := string("serve ip is " + serve_url + "\n")

		// 表单获取以及返回输出
		fmt.Println(r.Form)
		url_string := strings.Join(r.Form["IPnum"], "")
		fmt.Println("ip:", url_string)
		if PingTest(url_string) {
			result = result + "ping " + url_string + " success!"
		} else {
			result = result + "ping " + url_string + " failed!"
		}

		fmt.Fprint(w, result)
	}
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53") //建立连接
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return ip, err
}
