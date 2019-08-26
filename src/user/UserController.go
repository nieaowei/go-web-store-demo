/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :
 *******************************************************/
package user

import (
	"encoding/json"
	"fmt"
	"go-web-store-demo/src/commons"
	"net/http"
)

func UserHandler() {
	commons.MainRouter.HandleFunc("/login", loginController)
	commons.MainRouter.HandleFunc("/register", registerController)
}

func loginController(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)
	er := loginService(username, password)
	data, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(data)
}

func registerController(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	fmt.Println(username, password, email, phone)
}
