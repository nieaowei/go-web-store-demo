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
	data := selectByPage(rows, page)
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

func getTotalService() (res commons.Result) {
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

func addItemByItemService(item *TbItem) (res commons.Result) {
	item.ID = generateItemId()
	lenth := addByItem(item)
	res.Status = 400
	if lenth != 0 {
		//增加失败
		return
	}
	res.Status = 200
	return
}
