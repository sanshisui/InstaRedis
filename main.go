package main

import (
	"InstaRedis/handle/conn"
	"fmt"
	"net/http"
)


func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应内容类型
	w.Header().Set("Content-Type", "application/json")
	// 向客户端发起响应
	fmt.Fprint(w, `{"message": "Hello, World!"}`)
}



func main() {
	// 设置路由和处理器
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/conn", conn.ConnHandler)

	// 启动web服务器,监听8080端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器启动失败", err)
	}

}


