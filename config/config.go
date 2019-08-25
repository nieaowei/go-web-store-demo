/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :
 *******************************************************/
package config

import (
	"encoding/json"
	"io/ioutil"
)

var (
	config []byte
)

func init() {
	err := readConfig()
	if err != nil {
		//@todo
		panic(err)
	}
}

//读取配置文件
func readConfig() (err error) {
	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		//@todo
		return
	}
	config = data
	return
}

//获取key相关配置的map
func GetConfigData(key string) (res map[string]interface{}) {
	temp := make(map[string]interface{})
	err := json.Unmarshal(config, &temp)
	if err != nil {
		//@todo
		return
	}
	return temp[key].(map[string]interface{})
}
