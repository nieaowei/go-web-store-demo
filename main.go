/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :
 *******************************************************/
package main

import (
	"go-web-store-demo/src/commons"
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
	commons.MainRouter.HandleFunc("/", login)
	commons.MainRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	commons.MainRouter.PathPrefix("/view/").Handler(http.StripPrefix("/view/", http.FileServer(http.Dir("view"))))
	user.UserHandler()
	http.ListenAndServe(":8080", commons.MainRouter)
}
