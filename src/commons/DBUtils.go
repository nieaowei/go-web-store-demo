/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :
 *******************************************************/
package commons

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"go-web-store-demo/config"
)

type MyDataBase struct {
	dB         *sql.DB
	stmt       *sql.Stmt
	rows       *sql.Rows
	Driver     string `json:"driver"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Table      string `json:"table"`
	DataSource string `json:"dataSource"`
}

var (
	MyDB = new(MyDataBase)
)

//默认数据库初始化
func init() {
	err := dataBaseConfigInit()
	if err != nil {
		//@todo
		return
	}
}

//默认数据库初始化
func dataBaseConfigInit() (err error) {
	data, err := json.Marshal(config.GetConfigData("database"))
	if err != nil {
		//@todo
		return
	}
	err = json.Unmarshal(data, MyDB)
	if err != nil {
		return
	}
	return
}

//打开数据库链接
func (this *MyDataBase) openConn() (err error) {
	this.dB, err = sql.Open("mysql", this.Name+":"+this.Password+"@tcp("+this.DataSource+")/"+this.Table)
	if err != nil {
		//@TODO 实现日志文件记录
		return
	}
	return nil
}

//关闭数据库链接
func (this *MyDataBase) CloseConn() {
	if this.rows != nil {
		this.rows.Close()
	}
	if this.stmt != nil {
		this.stmt.Close()
	}
	if this.dB != nil {
		this.dB.Close()
	}

}

//
func (this *MyDataBase) Dml(sql string, args ...interface{}) (lenth int64, err error) {
	err = this.openConn()
	if err != nil {
		//@TODO 日志
		return
	}
	this.stmt, err = this.dB.Prepare(sql)
	if err != nil {
		//@TODO
		return
	}
	result, err := this.stmt.Exec(args...)
	if err != nil {
		//@TODO
		return
	}
	lenth, err = result.RowsAffected()
	if err != nil {
		//@todo
		return
	}
	this.CloseConn()
	return
}

func (this *MyDataBase) Dql(sql string, args ...interface{}) (rows *sql.Rows, err error) {
	err = this.openConn()
	if err != nil {
		//@TODO 日志
		return
	}
	this.stmt, err = this.dB.Prepare(sql)
	if err != nil {
		//@TODO
		return
	}
	rows, err = this.stmt.Query(args...)
	if err != nil {
		//@TODO
		return
	}
	//this.CloseConn()//调用此函数要记得关闭
	return
}
