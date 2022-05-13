package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"myproject/common"
	"myproject/router"
	"os"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = router.CollectRouter(r)

	r.Run(":9090")
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")            //设置读取的文件名
	viper.SetConfigType("yml")               //设置读取的文件类型
	viper.AddConfigPath(workDir + "/config") //设置读取的文件路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}
