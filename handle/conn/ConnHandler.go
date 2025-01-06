package conn

import (
	"encoding/json"
	"net/http"
	"InstaRedis/obj"
)

// ConnHandler is a handler for connection
// ConnHandler 是一个redis连接处理器
func ConnHandler(w http.ResponseWriter, r *http.Request) {

	//只处理POST请求
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	var connInfo obj.ConnInfo

	json.NewDecoder(r.Body).Decode(&connInfo)


	// 设置响应内容类型
	w.Header().Set("Content-Type", "application/json")
	// 向客户端发起响应
	w.Write([]byte(`{"message": "ConnHandler 是一个redis连接处理器"}`))
}