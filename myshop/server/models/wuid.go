package models

import (
	"database/sql"
	"fmt"
	"github.com/edwingeng/wuid/mysql/wuid"
)
var g = wuid.NewWUID("default",nil)
func CreateUid(dsn string) {
	newDB := func() (*sql.DB, bool, error) {
		db,err := sql.Open("mysql",dsn)
		if err != nil {
			panic("create wuid err:"+ err.Error())
		}
		return db, true, nil
	}
	_ = g.LoadH28FromMysql(newDB,"wuid")
}
func GetWuid() string {
	return fmt.Sprintf("%#016x",g.Next())
}
