/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-30
 *  Notes       :
 *******************************************************/
package cat

import (
	"fmt"
	"go-web-store-demo/src/commons"
)

func selectByIdDao(id int) (res *TbItemCat) {
	rows, err := commons.MyDB.Dql("select *from tb_item_cat where id=?", id)
	if err != nil {
		//@todo
		fmt.Println(err)
		return
	}
	if rows.Next() {
		res = new(TbItemCat)
		rows.Scan(&res.Id, &res.ParentId, &res.Name, &res.Status, &res.SortOrder, &res.IsParent, &res.Created, &res.Updated)
	}
	commons.MyDB.CloseConn()
	return
}
