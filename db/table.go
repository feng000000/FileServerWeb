package db

import (
    "time"
)


type Model struct {
    ID uint `gorm:"primaryKey"`
}

type User struct {
    Model
    UUID        string      `json:"uuid"`
    Level       int8        `json:"level"`
    Banned      bool        `json:"banned"`
    Username    string      `json:"username"`
    Password    string      `json:"-"`
    Email       string      `json:"email"`
    Created     time.Time   `json:"created" gorm:"autoCreateTime"`
    LastLogin   time.Time   `json:"last_login" gorm:"autoCreateTime"`
}

type File struct {
    Model
    UUID        string      `json:"uuid"`
    Filename    string      `json:"filename"`
    Created     time.Time   `json:"created" gorm:"autoCreateTime"`
    Size_kB     int         `json:"size_kb"`
}

type LevelStorge struct {
    Model
    Level       int8        `json:"level"`
    StorgeLimit int         `json:"storge_limit"` // 单位kB
}