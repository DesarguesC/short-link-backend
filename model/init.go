package model

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	connectDatabase()
	err := DB.AutoMigrate(&Url{}, &Users{}) // TODO: add table structs here
	if err != nil {
		logrus.Fatal(err)
	}
}

func connectDatabase() {
	viper.SetConfigName("conf")
	//viper.SetConfigType("json")	yaml also be fine
	viper.AddConfigPath("./")
	// 数据库用户登录的配置文件直接放在short-link-backend文件夹内
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	loginInfo := viper.GetStringMapString("sql")
	dbArgs := loginInfo["username"] + ":" + loginInfo["password"] +
		"@(localhost)/" + loginInfo["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
}
