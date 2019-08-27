/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :
 *******************************************************/
package user

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
)
