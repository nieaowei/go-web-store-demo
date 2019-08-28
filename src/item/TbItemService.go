/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-26
 *  Notes       :
 *******************************************************/
package item

import (
	"fmt"
	"go-web-store-demo/src/commons"
)

func showItemSerive(page, row int) (res *commons.Result) {
	data := selectByPageDao(page, row)
	if data != nil { //查询到数据
		for i, v := range data {
			fmt.Println(i, v)
		}
		return
	}
	return nil
}
