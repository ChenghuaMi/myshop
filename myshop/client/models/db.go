package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myshop/client/config"
)

var DB *gorm.DB

const DSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4"
func InitDb(config *config.Config) {
	var err error
	username := config.Mysql.Username
	password := config.Mysql.Password
	host := config.Mysql.Host
	port := config.Mysql.Port
	db := config.Mysql.Db
	dsn := fmt.Sprintf(DSN,username,password,host,port,db)
	DB,err = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("db init error %s \n",err.Error()))
	}
	InitWuid(dsn)
	fmt.Println(GetWuid())
}
