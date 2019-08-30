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

func showItemSerive(rows, page int) (res commons.Result) {
	data := selectByPageDao(rows, page)
	if data != nil { //查询到数据
		for i, v := range data {
			fmt.Println(i, v)
		}
		res.Data = data
		res.Status = 200
	} else {
		res.Status = 400
	}
	return
}

func getTotal() (res commons.Result) {
	r, err := commons.MyDB.Dql("select count(*) from tb_item")
	if err != nil {
		fmt.Println(err)
		res.Status = 400
		return
	}
	res.Status = 200
	for r.Next() {
		r.Scan(&res.Data)
	}
	commons.MyDB.CloseConn()
	return
}
