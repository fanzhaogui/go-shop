package main

import (
	"net/http"
	"html/template"
	"shop/src/user"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/login.html")
	t.Execute(w, nil)
}

func main()  {
	s := http.Server{Addr: ":8088"}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// 类似 设置router
	http.HandleFunc("/", welcome)
	user.UserHandler()

	s.ListenAndServe()
}
