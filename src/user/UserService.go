/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :业务逻辑
 *******************************************************/
package user

import (
	"go-web-store-demo/src/commons"
)

func loginService(usern, pwd string) (res commons.Result) {
	user := SelectByPwd(usern, pwd)
	if user != nil { //查询到数据
		res.Status = 200
	} else { //没有数据
		res.Status = 400
	}
	return
}

func registerService(username, pwd, email, phone string) (res commons.Result) {

	return
}
