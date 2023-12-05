package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db * gorm.DB
)

func Connect() {
   var err error

  dsn := "root:@tcp(127.0.0.1:3306)/bookstoredb?charset=utf8mb4&parseTime=True&loc=Local"
  db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	
	fmt.Println("connected to database")
}

func GetDB() *gorm.DB{
	return db
}