package db

import (
	"fmt"
	"log"
	"time"

	"FileServerWeb/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

type Model struct {
	ID uint `gorm:"primaryKey"`
}

type User struct {
	Model
	UUID 		string
	Username	string
	Password	string
	Email		string
	Created    	time.Time
	LastLogin	time.Time
}


func init() {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_ADDR,
		config.DB_NAME,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(
		&User{},
	)
}
