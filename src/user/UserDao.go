/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :数据访问
 *******************************************************/
package user

import (
	"fmt"
	"go-web-store-demo/src/commons"
	"regexp"
	"time"
)

//通过密码去查询数据，查询出错返回nil，没有查询到数据返回nil
func selectByPwd(usern, pwd string) (user *User) {
	sql := "select * from tb_user where username=? and password=? or email=? and password=? or phone=? and password=?"
	rows, err := commons.MyDB.Dql(sql, usern, pwd, usern, pwd, usern, pwd)
	if err != nil {
		//@todo
		return
	}
	if rows.Next() {
		user = new(User)
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Phone, &user.Email, &user.Created, &user.Updated)
		if err != nil {
			//@todo 查询出错
			return nil
		}
		commons.MyDB.CloseConn()
		return
	}
	return
}

//通过用户名，密码，手机号码，电子邮箱去新增一条记录，并且逐级判断关键键值是否存在，由于
func addUserByUPP(u *User) int8 {
	/*
		下列操作效率低
	*/
	//sql := "select * from tb_user where username=?"
	//rows , err := commons.MyDB.Dql(sql,username)
	//if err != nil {
	//	//@todo
	//	commons.MyDB.CloseConn()
	//	return USERERROR
	//}
	//if rows.Next()==true{
	//	commons.MyDB.CloseConn()
	//	return EXIST_USERNAME
	//}
	//sql = "select * from tb_user where phone=?"
	//rows , err = commons.MyDB.Dql(sql,phone)
	//if err !=nil{
	//	commons.MyDB.CloseConn()
	//	return USERERROR
	//}
	//if rows.Next()==true{
	//	commons.MyDB.CloseConn()
	//	return EXIST_PHONE
	//}
	//sql = "select * from tb_user where email=?"
	//rows , err = commons.MyDB.Dql(sql,email)
	//if err !=nil{
	//	commons.MyDB.CloseConn()
	//	return USERERROR
	//}
	//if rows.Next()==true{
	//	commons.MyDB.CloseConn()
	//	return EXIST_EMAIL
	//}
	//如果存在就无法执行下面语句
	//以下通过分析错误信息，简化了代码
	sql := "insert into tb_user values(default,?,?,?,?,?,?)"
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	lenth, err := commons.MyDB.Dml(sql, u.Username, u.Password, u.Phone, u.Email, timeStr, timeStr)
	if err != nil { //解析错误
		r, _ := regexp.Compile("'[a-zA-Z0-9]+'")
		temp := r.FindAllString(fmt.Sprintf("%s", err), 2)
		switch temp[len(temp)-1] { //在这里有一个小坑，不能直接去判断temp[2]，直接判断会产生恐慌，所以需要取长度取进行判断
		case "'username'":
			return EXIST_USERNAME
		case "'phone'":
			return EXIST_PHONE
		case "'email'":
			return EXIST_EMAIL
		default:
			return USERERROR
		}
	}
	if lenth < 0 {
		return USERERROR
	}
	return USEROK
}
