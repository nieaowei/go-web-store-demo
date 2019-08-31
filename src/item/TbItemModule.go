/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-26
 *  Notes       :
 *******************************************************/
package item

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
