/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-26
 *  Notes       :
 *******************************************************/
package item

type TbItem struct {
	ID         int64
	Title      string
	Sell_point string
	Price      int64
	Num        int64
	Barcode    string
	Image      string
	Cid        int
	Status     int8
	Created    string
	Updated    string
}
