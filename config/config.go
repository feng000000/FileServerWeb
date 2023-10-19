package config

import (
	"github.com/pelletier/go-toml"
)

var (
	config, _ = toml.LoadFile("./config.toml")

	HOME_DIR = config.Get("HOME_DIR").(string)
	FILE_PATH = config.Get("FILE.FILE_PATH").(string)
	SECRET_KEY = []byte(config.Get("SECRET_KEY").(string))

)