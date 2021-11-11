package main

import (
	"log"
	"net/http"
	"webtest"
)

func main() {
	http.HandleFunc("/", webtest.SayhelloName) //设置访问的路由
	http.HandleFunc("/login", webtest.Login)
	errs := http.ListenAndServe(":9090", nil)
	if errs != nil {
		log.Fatal("ListenAndServe: ", errs)
	}

}
