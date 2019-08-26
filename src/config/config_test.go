/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :
 *******************************************************/
package config

import (
	"reflect"
	"testing"
)

func TestGetConfigData(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantRes map[string]interface{}
	}{
		// TODO: Add test cases.
		{
			"1",
			args{key: "database"},
			map[string]interface{}{"driver": "mysql", "name": "root", "password": "nieaowei", "table": "ego", "dataSource": "localhost:3306"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetConfigData(tt.args.key); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetConfigData() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
