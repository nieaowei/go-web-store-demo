/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-26
 *  Notes       :
 *******************************************************/
package item

import (
	"go-web-store-demo/src/commons"
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
				{
					4,
					"iphone4",
					"weqioii",
					1000,
					100,
					"",
					"",
					13,
					1,
					"2019-08-28 12:36:37",
					"2019-08-28 12:36:43",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotData := selectByPage(tt.args.rows, tt.args.page); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("selectByPage() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_getTotal(t *testing.T) {
	tests := []struct {
		name    string
		wantRes commons.Result
	}{
		// TODO: Add test cases.
		{
			"获取总条数",
			commons.Result{
				Status: 200,
				Data:   "13",
				Msg:    "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := getTotal(); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("getTotal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_addByItem(t *testing.T) {
	type args struct {
		item TbItem
	}
	tests := []struct {
		name    string
		args    args
		wantRes commons.Result
	}{
		// TODO: Add test cases.
		{
			"增加一条数据测试",
			args{item: TbItem{
				ID:         100,
				Title:      "iohon",
				Sell_point: "wads",
				Price:      100,
				Num:        10,
				Barcode:    "",
				Image:      "",
				Cid:        0,
				Status:     0,
				Created:    "",
				Updated:    "",
			}},
			commons.Result{
				Status: 200,
				Data:   nil,
				Msg:    "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := addByItem(tt.args.item); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("addByItem() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_alterById(t *testing.T) {
	type args struct {
		id    interface{}
		key   string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
	}{
		// TODO: Add test cases.
		{
			"更改状态测试",
			args{
				id:    4,
				key:   "status",
				value: 3,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := alterById(tt.args.id, tt.args.key, tt.args.value); gotRes != tt.wantRes {
				t.Errorf("alterById() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
