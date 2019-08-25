/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :数据访问
 *******************************************************/
package user

import "go-web-store-demo/src/commons"

func SelectByPwd(usern, pwd string) (user *User) {
	sql := "selcet * from tb_user where username=? and password=? or email=? and password=? or phone=? and password=?"
	rows, err := commons.MyDB.Dql(sql)
	if err != nil {
		//@todo
		return
	}
	if rows.Next() {
		user := new(User)
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Phone, &user.Email, &user.Created, &user.Updated)
		commons.MyDB.CloseConn()
		return
	}
	return nil
}
