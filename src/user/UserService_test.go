/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-27
 *  Notes       :
 *******************************************************/
package user

import (
	"go-web-store-demo/src/commons"
	"reflect"
	"testing"
)

func Test_matchEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantRes bool
	}{
		// TODO: Add test cases.
		{
			"邮箱匹配测试",
			args{email: "nieaowei@gmail.com"},
			true,
		},
		{
			"三级域名匹配测试",
			args{email: "nieaowei@vip.qq.com"},
			true,
		},
		{
			"后缀匹配测试",
			args{email: "nieaowei@vip..qq.com"},
			false,
		},
		{
			"多限定符匹配测试",
			args{email: "nieaowei@@qq.com"},
			false,
		},
		{
			"四级域名匹配测试",
			args{email: "nieaowei@123.vip.qq.com"},
			true,
		},
		{
			"匹配测试",
			args{email: "nieaowei123@qq.com"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := matchEmail(tt.args.email); gotRes != tt.wantRes {
				t.Errorf("matchEmail() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_matchPhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name    string
		args    args
		wantRes bool
	}{
		// TODO: Add test cases.
		{
			"长度匹配测试",
			args{phone: "1357513257"},
			false,
		},
		{
			"正确匹配测试",
			args{"13575132755"},
			true,
		},
		{
			"前缀匹配测试",
			args{phone: "23356567676"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := matchPhone(tt.args.phone); gotRes != tt.wantRes {
				t.Errorf("matchPhone() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_loginService(t *testing.T) {
	type args struct {
		usern string
		pwd   string
	}
	tests := []struct {
		name    string
		args    args
		wantRes commons.Result
	}{
		// TODO: Add test cases.
		{
			"成功登陆测试",
			args{
				usern: "nieaowei",
				pwd:   "nieaowei",
			},
			commons.Result{
				Status: 200,
				Data:   nil,
				Msg:    "",
			},
		},
		{
			"账号错误测试",
			args{
				usern: "1242134",
				pwd:   "nieaowei",
			},
			commons.Result{
				Status: 400,
				Data:   nil,
				Msg:    "",
			},
		},
		{
			"密码错误测试",
			args{
				usern: "nieaowei",
				pwd:   "12313",
			},
			commons.Result{
				Status: 400,
				Data:   nil,
				Msg:    "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := loginService(tt.args.usern, tt.args.pwd); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("loginService() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
