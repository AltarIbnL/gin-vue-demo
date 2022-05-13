package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"myproject/Model"
	"net/url"
)

var Db *gorm.DB

func InitDB() *gorm.DB {
	driver := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	passwd := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		passwd,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	db, err := gorm.Open(driver, args)

	if err != nil {
		panic("fail to connect to database,err:" + err.Error())
	}
	//自动根据表的类型创建数据表
	db.AutoMigrate(&Model.User{})
	Db = db // 赋值给下面的GetDB
	return db
}

func GetDB() *gorm.DB {
	return Db
}
