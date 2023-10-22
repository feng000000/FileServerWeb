package config

import (
	"log"

	"github.com/pelletier/go-toml"
)

var (
	config, err = toml.LoadFile("./config.toml")


	HOME_DIR = config.Get("HOME_DIR").(string)
	FILE_PATH = config.Get("FILE.FILE_PATH").(string)
	SECRET_KEY = []byte(config.Get("SECRET_KEY").(string))

	DB_USERNAME = config.Get("DATABASE.USERNAME").(string)
	DB_PASSWORD = config.Get("DATABASE.PASSWORD").(string)
	DB_ADDR = config.Get("DATABASE.ADDR").(string)
	DB_NAME = config.Get("DATABASE.DB_NAME").(string)
)

func init(){
	var err error
	var config *toml.Tree
	config, err = toml.LoadFile("./config.toml")

	if err != nil {
		log.Fatal(err)
	}

	HOME_DIR = config.Get("HOME_DIR").(string)
}