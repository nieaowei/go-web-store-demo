/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-27
 *  Notes       :
 *******************************************************/
package user

import "testing"

func Test_addUserByUPP(t *testing.T) {
	type args struct {
		username string
		password string
		phone    string
		email    string
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		// TODO: Add test cases.
		{
			"用户名存在测试",
			args{
				username: "nieaowei",
				password: "nieaowei",
				phone:    "13575132575",
				email:    "nieaowei@qq.com",
			},
			EXIST_USERNAME,
		},
		{
			"手机号码存在测试",
			args{
				username: "",
				password: "nieaowei",
				phone:    "13575132575",
				email:    "nieaowei@qq.com",
			},
			EXIST_PHONE,
		},
		{
			"邮箱存在测试",
			args{
				username: "",
				password: "nieaowei",
				phone:    "",
				email:    "nieaowei@qq.com",
			},
			EXIST_EMAIL,
		},
		{
			"新增用户测试",
			args{
				username: "nieaowei12",
				password: "nieaowei12",
				phone:    "13575132576",
				email:    "nieaowei11@qq.com",
			},
			USEROK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addUserByUPP(tt.args.username, tt.args.password, tt.args.phone, tt.args.email); got != tt.want {
				t.Errorf("addUserByUPP() = %v, want %v", got, tt.want)
			}
		})
	}
}
