package common

import (
	"fmt"
	"gin_project/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "gin_db"
	username := "root"
	password := "1115774750"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	return db
}

func GetDb() *gorm.DB {
	return InitDatabase()
}
