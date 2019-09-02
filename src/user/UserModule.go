/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :
 *******************************************************/
package user

import (
	"encoding/json"
	"net/http"
	"strings"
)

//对应数据库的用户表
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

/*
Some return values for the registration function.
*/
const (
	USEROK = iota
	USERERROR
	EXIST_USERNAME
	EXIST_PHONE
	EXIST_EMAIL
	DATA_USERNAME_ERROR
	DATA_PHONE_ERROR
	DATA_EMAIL_ERROR
)

/*
	To create a User object by request, I don't know the efficiency of this method for the time being.
*/
func NewUserByRequest(r *http.Request) (user *User) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	user = new(User)
	data, _ := json.Marshal(r.Form)
	tempStr := strings.ReplaceAll(string(data), "[", "")
	tempStr = strings.ReplaceAll(tempStr, "]", "")
	err = json.Unmarshal([]byte(tempStr), user)
	if err != nil {
		return nil
	}
	// The following code has relized in other function.
	//if !user.MatchPhone(){
	//	return nil
	//}
	//if !user.MatchEmail(){
	//	return nil
	//}
	return
}
