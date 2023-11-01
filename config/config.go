package config

import (
	"log"

	"github.com/pelletier/go-toml"
	"path/filepath"
	// "reflect"
)

var (
	HOME_DIR		string
	DATA_FILE_PATH	string
	LOG_FILE_PATH	string
	SECRET_KEY		[]byte
	DB_USERNAME		string
	DB_PASSWORD		string
	DB_ADDR			string
	DB_NAME			string
)

func init() {
	var err error
	var config *toml.Tree
	config, err = toml.LoadFile("./config.toml")

	if err != nil {
		log.Fatal(err)
	}

	HOME_DIR 		= config.Get("HOME_DIR").(string)
	SECRET_KEY 		= []byte(config.Get("SECRET_KEY").(string))

	// FILE
	DATA_FILE_PATH  = parsePath(config.Get("FILE.DATA_FILE_PATH").(string))
	LOG_FILE_PATH  	= parsePath(config.Get("FILE.LOG_FILE_PATH").(string))

	// DB
	DB_USERNAME		= config.Get("DATABASE.USERNAME").(string)
	DB_PASSWORD 	= config.Get("DATABASE.PASSWORD").(string)
	DB_ADDR 		= config.Get("DATABASE.ADDR").(string)
	DB_NAME 		= config.Get("DATABASE.DB_NAME").(string)
}

func parsePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(HOME_DIR, path)
}
