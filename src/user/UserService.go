/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :业务逻辑
 *******************************************************/
package user

import (
	"go-web-store-demo/src/commons"
	"regexp"
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

//@todo 需要去实现，用户名，邮箱，手机号码存在的响应，暂时搁置
func registerService(username, pwd, email, phone string) (res commons.Result) {
	//首先验证是否符合要求
	if !matchEmail(email) {
		res.Status = 400
		return
	}
	if !matchPhone(phone) {
		res.Status = 200
		return
	}
	//如果都符合要求继续执行注册
	err := addUserByUPP(username, pwd, email, phone)
	if err != 0 {
		res.Status = 400
		return
	} else {
		res.Status = 200
		return
	}
}

//以下作为数据验证判断功能
func matchEmail(email string) (res bool) {
	res, _ = regexp.MatchString(`^([\w\.\_]{2,})@(\w{1,})(\.[\w]+)+$`, email)
	return
}

func matchPhone(phone string) (res bool) {
	res, _ = regexp.MatchString(`^(1[3|4|5|6|7|8|9][0-9][\d]{8})$`, phone)
	return
}
