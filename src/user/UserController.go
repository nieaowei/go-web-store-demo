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
	user := User{
		ID:       0,
		Username: username,
		Password: password,
		Phone:    username,
		Email:    username,
		Created:  "",
		Updated:  "",
	}
	res := user.LoginVerify()
	data, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(data)
}

func registerController(w http.ResponseWriter, r *http.Request) {
	user := NewUserByRequest(r) //不知道效率
	res := user.RegisterVerify()
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(data)
}
