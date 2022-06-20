package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	upgrade()
}

// 向服务器发送更新消息
func upgrade() {
	addr := flag.String("sa", "", "please input your server's api address of upgrade")
	flag.Parse()
	client := http.DefaultClient
	resp, err := client.Get(*addr)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
