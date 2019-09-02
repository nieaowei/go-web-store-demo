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

/*
Query data in the database by account number and password .
Return nil if the query has an error or the data does not exist.
*/
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

/*
Add a record to the database based on the user object and be able to analyze the error.
According to the return value , you can know the cause of the error and can handle it
according to different reasons
*/
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
	//If result is existed, the following code will be not run.

	// The following code implements the function of error location bt analyzing the error information.
	//And improve operational efficiency.
	sql := "insert into tb_user values(default,?,?,?,?,?,?)"
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	lenth, err := commons.MyDB.Dml(sql, u.Username, u.Password, u.Phone, u.Email, timeStr, timeStr)
	if err != nil { //解析错误
		r, _ := regexp.Compile("'[a-zA-Z0-9]+'")
		temp := r.FindAllString(fmt.Sprintf("%s", err), 2)
		switch temp[len(temp)-1] {
		// There are a point in here,you should not judge temp[2] directly,otherwise result in panic.
		// The solution is to take the judgment by acquiring the length of the data slice.
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
