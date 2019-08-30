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

//注册时使用,进行数据查询判断是否存在，返回数据
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
通过请求的参数去创建用户对象,暂时不知道该方式的效率
//r.ParseForm()
//qq,_:= json.Marshal(r.Form)
//tempStr:=strings.ReplaceAll(string(qq),"[","")
//tempStr = strings.ReplaceAll(tempStr,"]","")
//fmt.Println(tempStr)
//user := User{}
//json.Unmarshal([]byte(tempStr),&user)
//fmt.Println(user)
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
	//if !user.MatchPhone(){
	//	return nil
	//}
	//if !user.MatchEmail(){
	//	return nil
	//}
	return
}
