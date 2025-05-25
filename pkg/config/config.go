package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

var (
	db *gorm.DB
)

const (
	port = 3306
	user = "root"
	pass = "Mir@ge308"
	dbname = "Amitis"
)

// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

func Connection() *gorm.DB{
	//sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, port, dbname)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Error in DB connection: %v", err))
	}

	fmt.Println("DB is Connected!")
	return db
}

