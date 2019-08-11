package user

import (
	"net/http"
	"encoding/json"
)

func UserHandler()  {
	http.HandleFunc("/login", loginController)
}

// 登录
func loginController(w http.ResponseWriter, r *http.Request)  {
	// 接收参数
	// 验证
	// 响应
	un := r.FormValue("username")
	pwd := r.FormValue("password")
	re := LoginService(un, pwd)
	// 把结构体转为json数据
	b, _ := json.Marshal(re)

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}