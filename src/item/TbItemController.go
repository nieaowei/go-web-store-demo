/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-26
 *  Notes       :
 *******************************************************/
package item

import (
	"encoding/json"
	"go-web-store-demo/src/commons"
	"net/http"
	"strconv"
)

func TbItemHandler() {
	commons.MainRouter.HandleFunc("/showItem", showItemController)
	commons.MainRouter.HandleFunc("/getTotal", getTotalController)
	commons.MainRouter.HandleFunc("/addItem", addItemByItemController)
}

func showItemController(w http.ResponseWriter, r *http.Request) {
	rows, _ := strconv.Atoi(r.FormValue("rows"))
	page, _ := strconv.Atoi(r.FormValue("page"))
	res := showItemSerive(rows, page)
	data, _ := json.Marshal(res)
	w.Write(data)
}

func getTotalController(w http.ResponseWriter, r *http.Request) {
	res := getTotalService()
	data, _ := json.Marshal(res)
	w.Write(data)
}

func addItemByItemController(w http.ResponseWriter, r *http.Request) {
	item := NewTbItemByRequest(r)
	res := addItemByItemService(item)
	data, _ := json.Marshal(res)
	w.Write(data)
}
