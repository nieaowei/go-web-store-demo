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
			"以2行为一页查询第一页的数据",
			args{
				rows: 2,
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
			},
				{
					2,
					"test1",
					"ttttt",
					12000,
					100,
					"",
					"",
					11,
					1,
					"2019-08-26 23:17:10",
					"2019-08-26 23:17:16",
				},
			},
		},
		{
			"以2行为一页查询第二页的数据",
			args{
				rows: 2,
				page: 2,
			},
			[]TbItem{
				{
					3,
					"qwe",
					"qwewqeq",
					123,
					321,
					"",
					"",
					12,
					1,
					"2019-08-27 23:10:19",
					"2019-08-27 23:10:20",
				},
			},
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
