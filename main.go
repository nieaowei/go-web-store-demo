/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :
 *******************************************************/
package main

import (
	"go-web-store-demo/config"
	_ "go-web-store-demo/src/commons"
	"go-web-store-demo/src/user"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/login.html")
	t.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/register.html")
	t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr: config.GetConfigData("server")["addr"].(string),
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", login)
	user.UserHandler()
	server.ListenAndServe()
}
