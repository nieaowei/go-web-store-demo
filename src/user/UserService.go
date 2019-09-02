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
	user := selectByPwd(usern, pwd)
	if user != nil { //查询到数据
		res.Status = 200
	} else { //没有数据
		res.Status = 400
	}
	return
}

//@todo 需要去实现，用户名，邮箱，手机号码存在的响应，暂时搁置
//func registerService(username, pwd, phone,email string) (res commons.Result) {
//	//首先验证是否符合要求
//	if !matchPhone(phone) {
//		res.Status = 200
//		return
//	}
//	if !matchEmail(email) {
//		res.Status = 400
//		return
//	}
//
//	//如果都符合要求继续执行注册
//	err := addUserByUPP(username, pwd, phone,email)
//	if err != 0 {
//		res.Status = 400
//		return
//	} else {
//		res.Status = 200
//		return
//	}
//}

//
func matchEmail(email string) (res bool) {
	res, _ = regexp.MatchString(`^([\w\.\_]{2,})@(\w{1,})(\.[\w]+)+$`, email)
	return
}

func matchPhone(phone string) (res bool) {
	res, _ = regexp.MatchString(`^(1[3|4|5|6|7|8|9][0-9][\d]{8})$`, phone)
	return
}

/*
	To verify that the user's account and password match.
*/
func (u *User) LoginVerify() (res commons.Result) {
	res.Status = 200
	switch {
	case u.Username != "":
		user := selectByPwd(u.Username, u.Password)
		if user != nil {
			return
		}
	case u.Phone != "":
		user := selectByPwd(u.Username, u.Password)
		if user != nil {
			return
		}
	case u.Email != "":
		user := selectByPwd(u.Username, u.Password)
		if user != nil {
			return
		}
	}
	res.Status = 400
	return
}

/*
	Format matching the mailbox and mobile phone number ,and call the database opration to add a user.
*/
func (u *User) RegisterVerify() (res commons.Result) {
	if !u.MatchEmail() {
		res.Status = DATA_EMAIL_ERROR
		res.Msg = "邮箱格式不正确"
		return
	}
	if !u.MatchPhone() {
		res.Status = DATA_PHONE_ERROR
		res.Msg = "手机号码格式不正确"
		return
	}
	if res1 := addUserByUPP(u); res1 != USEROK {
		res.Status = int(res1)
		switch res1 {
		case EXIST_USERNAME:
			res.Msg = "用户名已经存在"
		case EXIST_EMAIL:
			res.Msg = "邮箱已经存在"
		case EXIST_PHONE:
			res.Msg = "手机号码已经存在"
		}
		return
	}
	res.Status = 200
	return
}

/*
	Use regular expressions to match whether the data satisfies the format.
*/
func (u *User) MatchEmail() (res bool) {
	res, _ = regexp.MatchString(`^([\w\.\_]{2,})@(\w{1,})(\.[\w]+)+$`, u.Email)
	return
}
func (u *User) MatchPhone() (res bool) {
	res, _ = regexp.MatchString(`^(1[3|4|5|6|7|8|9][0-9][\d]{8})$`, u.Phone)
	return
}
