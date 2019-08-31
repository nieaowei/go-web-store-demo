/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-30
 *  Notes       :
 *******************************************************/
package cat

type TbItemCat struct {
	Id        int
	ParentId  int
	Name      string
	Status    byte
	SortOrder int8
	IsParent  byte
	Created   string
	Updated   string
}
