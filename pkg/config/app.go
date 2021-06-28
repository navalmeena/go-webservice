package config

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	// "fmt"
	"github.com/joho/godotenv"
)

var (
	db * gorm.DB
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
	    log.Fatal("Error loading .env file")
	 }

    
    host := os.Getenv("MYSQL_HOST")
    dbname := os.Getenv("MYSQL_DATABASE")
    user := os.Getenv("MYSQL_USERNAME")
    pswd := os.Getenv("MYSQL_PASSWORD")
    
    // fmt.Println("host: ",host)
    // fmt.Println("dbname: ",dbname)
    // fmt.Println("user: ", user)
    // fmt.Println("pswd: ",pswd)

	// Please define your user name and password for my sql.
	//d, err := gorm.Open("mysql", "root:droom@123@tcp(127.0.0.1:3306)/go-db?charset=utf8&parseTime=True&loc=Local")
	
	d, err := gorm.Open("mysql", user+":"+pswd+"@tcp("+host+":3306)/"+dbname+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		panic(err)
	}
	db = d
	db.LogMode(true)
	//db.SetLogger(gorm.Logger{revel.TRACE})
}

func GetDB() *gorm.DB {
	return db
}