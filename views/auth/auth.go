package auth

import (
    // "fmt"
    "bufio"
    "net/http"
    "os"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"

    "FileServerWeb/db"
    "FileServerWeb/widget/jwt"
    L "FileServerWeb/widget/logger"
    R "FileServerWeb/widget/response"
)


var DB = db.DB

func Login(c *gin.Context) {
    var err error
    var token = c.GetHeader("Authorization")

    if token != "" {
        _, err = jwt.ParseToken(token)
        if err == nil {
            c.JSON(http.StatusOK, R.Success(nil))
            return
        }
    }

    type LoginParams struct {
        Username    string `json:"username" binding:"required"`
        Password    string `json:"password" binding:"required"`
    }
    var param LoginParams
    if c.ShouldBind(&param) != nil {
        c.JSON(http.StatusBadRequest, R.BadRequest(nil))
        return
    }

    var user = db.User{
        Username: param.Username,
        Password: param.Password,
    }
    err = DB.Where(
        "username=? and password=?",
        param.Username,
        param.Password,
    ).Take(&user).Error
    if err != nil {
        c.JSON(http.StatusBadRequest, R.BadRequest(R.Json{
            "message": "Wrong username or password",
        }))
        return
    }

    var ret_token string
    ret_token, err = jwt.GenerateToken(param.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "InternalServerError",
        })
        return
    }
    c.JSON(http.StatusOK, R.Success(R.Json{
        "token" : ret_token,
    }))
    return
}


func Register(c *gin.Context) {
    var err error
    var code_file_path = "./activation_code.txt"

    // 参数判断
    type  RegisterParams struct {
        Username    string `json:"username" binding:"required"`
        Password    string `json:"password" binding:"required"`
        Code        string `json:"code" binding:"required"`
    }
    var param RegisterParams
    if c.ShouldBind(&param) != nil {
        c.JSON(http.StatusBadRequest, R.BadRequest(nil))
        return
    }

    // check activation code
    var file *os.File
    var codes = make(map[string]bool)
    file, err = os.OpenFile(code_file_path, os.O_RDWR|os.O_CREATE, 0666)
    defer file.Close()
    if err != nil {
        L.Logger.Error("file open failed")
    }
    defer file.Close()

    var scanner = bufio.NewScanner(file)
    for scanner.Scan() {
        var scan_code = strings.Split(scanner.Text(), " ")[0]
        var state = strings.Split(scanner.Text(), " ")[1]
        if state == "0" {
            codes[scan_code] = false
        } else {
            codes[scan_code] = true
        }
    }
    err = scanner.Err()
    if err != nil {
        L.Logger.Error(err.Error())
    }

    {
        var t bool
        var ok bool
        t, ok = codes[param.Code]
        if !ok || t != false {
            c.JSON(http.StatusBadRequest, R.BadRequest(R.Json{"message": "Invalid activation code"}))
            return
        }
    }

    result := DB.Where("Username = ?", param.Username).Find(&db.User{})
    if result.RowsAffected > 0 {
        c.JSON(http.StatusBadRequest, R.BadRequest(R.Json{"message": "Register failed"}))
        return
    }

    var user = db.User{
        UUID: uuid.NewString(),
        Level: 5,
        Banned: false,
        Username: param.Username,
        Password: param.Password,
    }
    DB.Create(&user)

    var ret_token string
    ret_token, err = jwt.GenerateToken(param.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "InternalServerError",
        })
        return
    }

    file.Close()
    file, err = os.OpenFile(code_file_path, os.O_WRONLY|os.O_TRUNC, 0666)
    defer file.Close()
    codes[param.Code] = true
    for k, v := range codes {
        if v {
            file.WriteString(k + " 1\n")
        } else {
            file.WriteString(k + " 0\n")
        }
    }

    c.JSON(http.StatusOK, R.Success(R.Json{"token": ret_token}))
    return
}