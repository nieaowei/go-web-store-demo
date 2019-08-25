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
	"net/http"
)

func UserHandler() {
	http.HandleFunc("/login", loginController)
}

func loginController(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(r)
	fmt.Println(username, password)
	er := LoginService(username, password)
	data, _ := json.Marshal(er)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(data)
}
