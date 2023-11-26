package db

import (
    "fmt"
    "log"

    "FileServerWeb/config"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

type Result *gorm.DB

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

    // 添加所有的数据库表
    DB.AutoMigrate(
        &User{},
        &File{},
        &LevelStorge{},
    )

    // 预添加数据
    var levelStorges = []LevelStorge{
        {Level: 0, StorgeLimit: (1 << 30) * 1},     // 管理员, 1TB
        {Level: 1, StorgeLimit: (1 << 30) * 1},     // 荣誉会员, 1TB
        {Level: 2, StorgeLimit: (1 << 20) * 500},   // 会员, 500GB
        {Level: 5, StorgeLimit: (1 << 20) * 10},   // 普通, 10GB
    }
    DB.Save(&levelStorges)
}
