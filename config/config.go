package config

import (
    "log"
    "os"

    "github.com/pelletier/go-toml"
    "path/filepath"
    // "reflect"
)

var (
    HOME_DIR            string
    USER_FILE_PATH      string
    LOG_FILE_PATH       string
    CODE_FILE_PATH      string
    SECRET_KEY          []byte
    DB_USERNAME         string
    DB_PASSWORD         string
    DB_ADDR             string
    DB_NAME             string
    EMAIL_SMTP_SERVER   string
    EMAIL_SMTP_PORT     int
    EMAIL_USERNAME      string
    EMAIL_PASSWORD      string
)

func init() {
    var err error
    var config *toml.Tree
    // config, err = toml.LoadFile("./config/config.toml")
    config, err = toml.LoadFile(os.Getenv("FILE_SERVER_CONFIG"))

    if err != nil {
        log.Fatal(err)
    }

    HOME_DIR    = config.Get("HOME_DIR").(string)
    SECRET_KEY  = []byte(config.Get("SECRET_KEY").(string))

    // FILE
    USER_FILE_PATH  = parsePath(config.Get("FILE.USER_FILE_PATH").(string))
    LOG_FILE_PATH   = parsePath(config.Get("FILE.LOG_FILE_PATH").(string))
    CODE_FILE_PATH  = parsePath(config.Get("FILE.CODE_FILE_PATH").(string))

    // DB
    DB_USERNAME = config.Get("DATABASE.USERNAME").(string)
    DB_PASSWORD = config.Get("DATABASE.PASSWORD").(string)
    DB_ADDR     = config.Get("DATABASE.ADDR").(string)
    DB_NAME     = config.Get("DATABASE.DB_NAME").(string)

    // EMAIL
    EMAIL_SMTP_SERVER   = config.Get("EMAIL.SMTP_SERVER").(string)
    EMAIL_SMTP_PORT     = int(config.Get("EMAIL.SMTP_PORT").(int64))
    EMAIL_USERNAME      = config.Get("EMAIL.USERNAME").(string)
    EMAIL_PASSWORD      = config.Get("EMAIL.PASSWORD").(string)
}

// 如果是绝对路径: 不变; 如果是相对路径: 和HOME_DIR拼接成绝对路径
func parsePath(path string) string {
    if filepath.IsAbs(path) {
        return path
    }
    return filepath.Join(HOME_DIR, path)
}