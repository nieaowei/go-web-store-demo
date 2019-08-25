/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :数据访问
 *******************************************************/
package user

import (
	"go-web-store-demo/src/commons"
)

//通过密码去查询数据，查询出错返回nil，没有查询到数据返回nil
func SelectByPwd(usern, pwd string) (user *User) {
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
