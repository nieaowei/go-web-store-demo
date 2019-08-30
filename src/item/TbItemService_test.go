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

func Test_showItemSerive(t *testing.T) {
	type args struct {
		page int
		row  int
	}
	tests := []struct {
		name    string
		args    args
		wantRes *commons.Result
	}{
		// TODO: Add test cases.
		{
			"以1行为1页查询第一页数据",
			args{
				page: 1,
				row:  1,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := showItemSerive(tt.args.page, tt.args.row); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("showItemSerive() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
