/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :
 *******************************************************/
package commons

//客户端服务端交互模板
type Result struct {
	Status int         `json:"status"` //状态为200成功
	Data   interface{} `json:"data"`   //返回数据
	Msg    string      `json:"msg"`    //返回信息
}
