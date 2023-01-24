package config

import(
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect(){
	d,err := gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/gobookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		log.Fatal(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}