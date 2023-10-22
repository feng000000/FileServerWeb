package db

import (
	"log"
	"database/sql"
    "github.com/jmoiron/sqlx"

	"github.com/go-sql-driver/mysql"

	"FileServerWeb/config"
)


var DB *sqlx.DB


type UsersTable struct {
	Username 	string
	Password 	string
	Email 		sql.NullString
}


func createTables() {
	var err error

	// TODO: 反射获取结构体的字段名, 然后创建数据库
	_, err = DB.Exec("create table if not exists users (username varchar(64), password varchar(64), email varchar(64));")
	if err != nil {
		log.Fatal(err)
	}
}


func init() {
	var cfg = mysql.Config{
		User:   config.DB_USERNAME,
		Passwd: config.DB_PASSWORD,
		Net:    "tcp",
		Addr:   config.DB_ADDR,
		DBName: config.DB_NAME,
        AllowNativePasswords: true,
	}

	var err error
	DB, err = sqlx.Open("mysql", cfg.FormatDSN())

	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	}

	createTables()
}
