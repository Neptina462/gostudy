package main

import (
	"flag"
	"fmt"
)

func main() {

	var arg1 string
	var arg2 string

	flag.StringVar(&arg1, "p", "hello", "rua")
	flag.StringVar(&arg2, "i", "haha", "pei")
	flag.Parse()

	fmt.Println(arg1)
	fmt.Println(arg2)

}
