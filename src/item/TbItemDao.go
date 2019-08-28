/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-26
 *  Notes       :
 *******************************************************/
package item

import (
	"database/sql"
	"fmt"
	"go-web-store-demo/src/commons"
)

/*
 */
func selectByPageDao(rows, page int) (data []TbItem) {
	r, err := commons.MyDB.Dql("select * from tb_item limit ?,?", rows*(page-1), rows*page)
	if err != nil {
		//@todo
		fmt.Println(err)
		return nil //没有查询到数据
	}
	for r.Next() {
		var t TbItem
		var s sql.NullString
		var s1 sql.NullString
		//如果直接使用t.barcode 由于数据库为null会导致填充失败
		r.Scan(&t.ID, &t.Title, &t.Sell_point, &t.Price, &t.Num, &s, &s1, &t.Cid, &t.Status, &t.Created, &t.Updated)
		t.Barcode = s.String
		t.Image = s1.String
		data = append(data, t)
	}
	commons.MyDB.CloseConn()
	return
}
