/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :业务逻辑
 *******************************************************/
package user

func LoginService(usern, pwd string) {
	user := SelectByPwd(usern, pwd)
	if user != nil {

	} else {
		//@todo 没有数据
		return
	}
}
