package main

import (
	"net/http"
	"html/template"
	"shop/src/user"
	"github.com/gorilla/mux"
	"shop/src/commons"
	"shop/src/item"
	"fmt"
	"shop/src/item/cat"
	"shop/src/item/param"
)

// 显示登入页面
func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/login.html")
	t.Execute(w, nil)
}

// restful 风格显示页面
func showPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("request vars ", vars)
	t, _ := template.ParseFiles("view/" + vars["page"] + ".html")
	t.Execute(w, nil)
}

func main() {
	// s := http.Server{Addr: ":8088"}
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//
	// // 类似 设置router
	// http.HandleFunc("/", welcome)
	// user.UserHandler()
	//
	// fmt.Println("127.0.0.1:8088")
	// s.ListenAndServe()

	commons.Router.HandleFunc("/", welcome)
	commons.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	commons.Router.HandleFunc("/page/{page}", showPage)

	// 调用所有user模块的handler
	user.UserHandler()
	// 商品
	item.ItemHandler()
	// 类目
	cat.CatHandler()

	param.TbItemParamHandler()

	// 规格参数

	// 商品规格参数

	fmt.Println("http service start: 127.0.0.1:8088")
	http.ListenAndServe(":8088", commons.Router)
}
