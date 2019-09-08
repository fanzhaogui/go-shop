package user

import (
	"net/http"
	"encoding/json"
	"shop/src/commons"
	"fmt"
)

func UserHandler()  {
	// http.HandleFunc("/login", loginController)
	commons.Router.HandleFunc("/login", loginController)
}

// 登录
func loginController(w http.ResponseWriter, r *http.Request)  {
	// 接收参数
	// 验证
	// 响应
	un := r.FormValue("username")
	pwd := r.FormValue("password")
	fmt.Println("do login: ", un, pwd)

	// 数据库密码需要加密？ 查询是否要携带上密码
	re := LoginService(un, pwd)
	// 把结构体转为json数据
	b, _ := json.Marshal(re)

	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}