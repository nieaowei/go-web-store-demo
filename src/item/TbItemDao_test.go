/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-26
 *  Notes       :
 *******************************************************/
package item

import (
	"reflect"
	"testing"
)

func Test_selectByPageDao(t *testing.T) {
	type args struct {
		rows int
		page int
	}
	tests := []struct {
		name     string
		args     args
		wantData []TbItem
	}{
		// TODO: Add test cases.
		{
			"测试一条数据",
			args{
				rows: 1,
				page: 1,
			},
			[]TbItem{{
				1,
				"test",
				"ttest",
				110,
				100,
				"",
				"www.baidu.com",
				10,
				1,
				"2019-08-26 22:49:45",
				"2019-08-26 22:49:49",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := selectByPageDao(tt.args.rows, tt.args.page); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("selectByPageDao() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
