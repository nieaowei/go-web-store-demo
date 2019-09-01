/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-26
 *  Notes       :
 *******************************************************/
package item

import (
	"encoding/json"
	"net/http"
	"strings"
)

type TbItem struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Sell_point string `json:"sell_point"`
	Price      int64  `json:"price"`
	Num        int64  `json:"num"`
	Barcode    string `json:"barcode"`
	Image      string `json:"image"`
	Cid        int    `json:"cid"`
	Status     int8   `json:"status"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
}

type TbItemChild struct {
	TbItem
	CategoryName string
}

//Returns nil if the operation fails,otherwise returns data successfully.
func NewTbItemByRequest(r *http.Request) (res *TbItem) {
	err := r.ParseForm()
	if err != nil {
		return nil
	}
	res = new(TbItem)
	data, err := json.Marshal(r.Form)
	tempStr := strings.ReplaceAll(string(data), "[", "")
	tempStr = strings.ReplaceAll(tempStr, "]", "")
	err = json.Unmarshal([]byte(tempStr), res)
	if err != nil {
		return nil
	}
	return
}
