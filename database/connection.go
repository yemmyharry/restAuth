package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//connect to the database
	dsn := "root:12345678@tcp(127.0.0.1:3306)/rest_auth?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = connection
	//connection.AutoMigrate(models.User{})

}
