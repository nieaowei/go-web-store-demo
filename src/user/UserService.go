/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :业务逻辑
 *******************************************************/
package user

import (
	"fmt"
	"go-web-store-demo/src/commons"
)

func LoginService(usern, pwd string) (res commons.Result) {
	user := SelectByPwd(usern, pwd)
	if user != nil { //查询到数据
		res.Status = 200
		fmt.Println("200")
	} else { //没有数据
		res.Status = 400
		fmt.Println("400")
	}
	return
}
