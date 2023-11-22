package file

import (
    "strings"
    "strconv"
    "errors"
    "net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "FileServerWeb/config"
    "FileServerWeb/db"
    "FileServerWeb/widget/jwt"
    L "FileServerWeb/widget/logger"
    R "FileServerWeb/widget/response"
)


var DB = db.DB


// 检查数据库中是否存在文件名
func checkFilename(filename string) (bool, error) {
    var result = DB.Model(&db.File{}).Where(
        "Filename = ?", filename,
    )

    if result.Error != nil {
        return false, result.Error
    } else if result.RowsAffected > 0 {
        return true, nil
    }
    return false, nil
}


func getAvailableFilename(filename string) (string, error) {
    var theFileName = filename
    var filenameExists, err = checkFilename(theFileName)
    if err != nil {
        return "", errors.New("InternalServerError")
    }

    var cnt = 1
    var strs = strings.Split(filename, ".")
    var _len = len(strs)
    var prefix = strs[0]
    for i := 1; i < _len - 1; i ++ {
        prefix = strs[i]
    }

    for filenameExists {
        theFileName = prefix + "_" + strconv.Itoa(cnt) + "." + strs[_len-1]

        filenameExists, err = checkFilename(theFileName)
        if err != nil {
            return "", errors.New("InternalServerError")
        }

        cnt ++
        if cnt >= 10 {
            return "", errors.New("Too many files with the same name")
        }
    }
    return theFileName, nil
}


// 获取用户的存储空间
func getStorgeUsage(UUID string) (int64, error) {
    var totalSize int64
    var files []db.File

    result := DB.Where("uuid = ?", "UUID").Find(&files)
    if result.Error != nil {
        return 0, result.Error
    }

    for _, file := range files {
        totalSize += file.Size_KB
    }

    return totalSize, nil
}


// 获取用户存储最大容量
func getStorgeLimit(UUID string) (int64, error) {
    var limit db.LevelStorge
    // 使用 Joins 连接两个表
    var result = DB.Joins("JOIN users ON users.level = level_storges.level").
                    Where("users.uuid = ?", UUID).
                    First(&limit)

    if result.Error != nil {
        return 0, result.Error
    }
    return limit.StorgeLimit, nil
}


// 表单上传文件, key为 "files"
func Upload(c *gin.Context) {
    var err error
    var result db.Result

    var token = c.GetHeader("Authorization")
    var claims *jwt.Claims
    if token != "" {
        claims, err = jwt.ParseToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, R.Unauthorized(nil))
            return
        }
    }

    var UUID = claims.UUID

    form, err := c.MultipartForm()
    if err != nil {
        L.Logger.Error(err.Error())
        c.JSON(http.StatusBadRequest, R.BadRequest(nil))
        return
    }

    // 获取所有文件
    files := form.File["files"]
    if len(files) == 0 {
        c.JSON(
            http.StatusBadRequest,
            R.BadRequest(R.Json{"message": "key [files] not found"}),
        )
        return
    }

    // 判断容量是否到达上限
    var allFileSize int64
    {
        var maxStorge int64
        var currentStorge int64

        maxStorge, err = getStorgeLimit(UUID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, R.DatabaseError(nil))
            return
        }
        currentStorge, err = getStorgeUsage(UUID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, R.DatabaseError(nil))
            return
        }

        for _, file := range files {
            allFileSize += file.Size / 10
        }

        // 可超 1G (1*1024*1024 kB)
        if currentStorge + allFileSize > maxStorge + int64(1 << 20) {
            c.JSON(
                http.StatusBadRequest,
                R.BadRequest(R.Json{"message":"Not enough storage space"}),
            )
            return
        }
    }

    // 存储文件
    for _, file := range files {
        // 检查数据库中是否存在文件名
        var theFileName, err = getAvailableFilename(file.Filename)
        if err != nil {
            c.JSON(http.StatusBadRequest, R.BadRequest(R.Json{"message":err.Error()}))
            return
        }

        // 在数据库中添加记录
        var new_file = db.File{
            UUID: UUID,
            Filename: theFileName,
            Size_KB: file.Size / 10,
        }

        result = DB.Create(&new_file)
        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, R.DatabaseError(nil))
        }

        var dst = filepath.Join(config.USER_FILE_PATH, UUID, theFileName)
        go c.SaveUploadedFile(file, dst)
    }

    c.JSON(http.StatusOK, R.Success(nil))
}


// 下载文件
func Download(c *gin.Context) {
    var err error
    var result db.Result

    var token = c.GetHeader("Authorization")
    var claims *jwt.Claims
    if token != "" {
        claims, err = jwt.ParseToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, R.Unauthorized(nil))
            return
        }
    }

    var UUID = claims.UUID

    type DownloadParams struct {
        FileName    string `json:"filename" binding:"required"`
    }
    var param DownloadParams
    if c.ShouldBind(&param) != nil {
        c.JSON(http.StatusBadRequest, R.BadRequest(nil))
        return
    }

    var file db.File
    result = DB.Where("filename = ?", param.FileName).First(&file)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        c.JSON(
            http.StatusBadRequest,
            R.BadRequest(R.Json{"message":"File dont exists"}),
        )
        return
    } else if result.Error != nil {
        c.JSON(http.StatusInternalServerError, R.DatabaseError(nil))
        return
    }

    var dst = filepath.Join(config.USER_FILE_PATH, UUID, file.Filename)
    c.File(dst)
}