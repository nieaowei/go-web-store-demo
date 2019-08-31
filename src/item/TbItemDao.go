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
	"time"
)

/*
以rows为一页查询第page页的数据
*/
func selectByPage(rows, page int) (data []TbItem) {
	r, err := commons.MyDB.Dql("select * from tb_item limit ?,?", rows*(page-1), rows)
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

func addByItem(item TbItem) (res commons.Result) {
	sql := "insert into tb_item values(?,?,?,?,?,?,?,?,default,?,?)"
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	r, err := commons.MyDB.Dml(sql, item.ID, item.Title, item.Sell_point, item.Price, item.Num, item.Barcode, item.Image, item.Cid, timeStr, timeStr)
	res.Status = 400
	if err != nil {
		//@todo
		fmt.Println(err)
		return
	}
	if r < 0 {
		return
	}
	res.Status = 200
	return
}

//通过id去更改相应的数据，返回-1失败  0也是失败  成功会返回成功的长度 1 如果更改的数据与
//原值相同返回0
func alterById(id interface{}, key string, value interface{}) (res int64) {
	sql := "update tb_item set "
	sql += key + "=? where id=?"
	fmt.Println(sql)
	res, err := commons.MyDB.Dml(sql, value, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
