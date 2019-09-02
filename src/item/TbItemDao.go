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
	"math/rand"
	"time"
)

//The first page of the record is queried  with a page line record as a page.
//Returns nil if no data is querired,otherwise returns the data slice of the query.
func selectByPage(rows, page int) (data []TbItem) {
	r, err := commons.MyDB.Dql("select * from tb_item limit ?,?", rows*(page-1), rows)
	if err != nil {
		//@todo
		fmt.Println(err)
		return nil //no data
	}
	for r.Next() {
		var t TbItem
		var s sql.NullString
		var s1 sql.NullString
		//Since the data in the database is null ,if you use t.barcide directly ,it will casuse the operation to fail.
		r.Scan(&t.ID, &t.Title, &t.Sell_point, &t.Price, &t.Num, &s, &s1, &t.Cid, &t.Status, &t.Created, &t.Updated)
		t.Barcode = s.String
		t.Image = s1.String
		data = append(data, t)
	}
	commons.MyDB.CloseConn()
	return
}

//Genereate a random number for use by the primary key

func generateItemId() (id int64) {
	rand.Seed(time.Now().UnixNano())
	id = rand.Int63n(10000000)
	fmt.Println(id)
	return
}

//Add a record to the table by *TbItem
func addByItem(item *TbItem) (res int) {
	sql := "insert into tb_item values(?,?,?,?,?,?,?,?,default,?,?)" //The default status is 1.  normal.
	timeStr := time.Now().Format("2006-01-02 15:04:05")              //Current time.
	r, err := commons.MyDB.Dml(sql, item.ID, item.Title, item.Sell_point, item.Price, item.Num, item.Barcode, item.Image, item.Cid, timeStr, timeStr)
	if err != nil {
		//@todo
		fmt.Println(err)
		return -1
	}
	if r < 0 {
		return -1
	}
	return
}

//The desired result is 1,but a result of 0 does not mean that the function is wrong,but the
//data that needs to be modified is the same as the data being modified
func alterById(id interface{}, key string, value interface{}) (res int64) {
	sql := "update tb_item set "
	sql += key + "=? where id=?"
	res, err := commons.MyDB.Dml(sql, value, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
