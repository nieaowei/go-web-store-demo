/*******************************************************
 *  ProjectName :go-web-store-demo
 *  Author      :nieaowei
 *  Date        :2019-08-24
 *  Notes       :
 *******************************************************/
package commons

import (
	"database/sql"
	"testing"
)

func TestMyDataBase_OpenConn(t *testing.T) {
	type fields struct {
		dB         *sql.DB
		stmt       *sql.Stmt
		rows       *sql.Rows
		Driver     string
		Name       string
		Password   string
		Table      string
		DataSource string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				dB:         new(sql.DB),
				stmt:       nil,
				rows:       nil,
				Driver:     "mysql",
				Name:       "root",
				Password:   "nieaowei",
				Table:      "ego",
				DataSource: "localhost:3306",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyDataBase{
				dB:         tt.fields.dB,
				stmt:       tt.fields.stmt,
				rows:       tt.fields.rows,
				Driver:     tt.fields.Driver,
				Name:       tt.fields.Name,
				Password:   tt.fields.Password,
				Table:      tt.fields.Table,
				DataSource: tt.fields.DataSource,
			}
			if err := this.openConn(); (err != nil) != tt.wantErr {
				t.Errorf("openConn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
