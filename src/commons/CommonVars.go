package commons

import "github.com/gorilla/mux"

// 路由控件
var (
	Router = mux.NewRouter()
	HEADER_CONTENT_TYPE string = "Content-Type"
	JSON_HEADER string = "application/json;charset=utf-8"
	CurrentPath string = "http://localhost:8088/"
)
