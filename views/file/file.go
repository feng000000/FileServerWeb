package file

import (
    "strings"
    "strconv"
    "errors"
    "net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"

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
func getStorgeUsage(username string) (int, error) {
    // TODO

    return 0, nil
}

// 增加存储空间
func addStorgeUsage(username string, storgeUsege int) error {
    // TODO

    return nil
}

func Test(c *gin.Context) {
    var user db.User
    var result = DB.Model(&user).Where("Username = ?", "username").First(&user)
    if result.Error != nil {
        println(result.Error.Error())
        c.JSON(http.StatusInternalServerError, R.DatabaseError(nil))
        return
    }

    c.JSON(http.StatusOK, R.Success(nil))
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

    var username = claims.Username

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
            R.BadRequest(R.Json{"message": "[files] not found"}),
        )
        return
    }

    for _, file := range files { // for files
        // 检查数据库中是否存在文件名
        var theFileName, err = getAvailableFilename(file.Filename)
        if err != nil {
            c.JSON(http.StatusBadRequest, R.BadRequest(R.Json{"message":err.Error()}))
            return
        }

        // TODO: 判断容量是否到达上限
        // var maxStorge int
        // var currentStorge int
        // var user db.User

        // // TODO: 多表查询获得用户最大容量
        // result = DB.Model(&User).Where("Username = ?", username).First(&user)
        // if result.Error != nil {
        //     c.JSON(http.StatusInternalServerError, R.DatabaseError(nil))
        //     return
        // }

        // currentStorge, err = getStorgeUsage(username)
        // if err != nil {
        //     c.JSON(http.StatusInternalServerError, R.DatabaseError(nil))
        //     return
        // }
        // if file.Size / 10 +

        // 在数据库中添加记录
        var user db.User
        result = DB.Model(&user).Where("Username = ?", username).First(&user)
        var UUID = user.Username
        var new_file = db.File{
            UUID: UUID,
            Filename: theFileName,
            Size_kB: int(file.Size / 10),
        }

        result = DB.Create(&new_file)
        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, R.DatabaseError(nil))
        }

        var dst = filepath.Join(config.USER_FILE_PATH, username, theFileName)
        go c.SaveUploadedFile(file, dst)
    } // for files end

    c.JSON(http.StatusOK, R.Success(nil))
}


func Download(c *gin.Context) {

}