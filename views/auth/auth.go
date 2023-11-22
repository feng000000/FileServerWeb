package auth

import (
	// "fmt"
	"bufio"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"FileServerWeb/db"
	"FileServerWeb/config"
	"FileServerWeb/widget/jwt"
	L "FileServerWeb/widget/logger"
	R "FileServerWeb/widget/response"
)


var DB = db.DB


func checkActivationCode(code string) (map[string]bool, bool) {
    var err error
    var file *os.File
    var codes = make(map[string]bool)

    file, err = os.OpenFile(
        config.CODE_FILE_PATH,
        os.O_RDWR|os.O_CREATE,
        0666,
    )
    defer file.Close()
    if err != nil {
        L.Logger.Error("file open failed")
        return nil, false
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

    if scanner.Err() != nil {
        L.Logger.Error(err.Error())
        return nil, false
    }

    var t bool
    var ok bool
    t, ok = codes[code]
    if !ok || t != false {
        return nil, false
    }
    return codes, true
}


func useActivationCode(codes map[string]bool, code string) bool {
    var err error
    var w_file *os.File
    w_file, err = os.OpenFile(
        config.CODE_FILE_PATH,
        os.O_WRONLY|os.O_TRUNC,
        0666,
    )
    if err != nil {
        return false
    }
    defer w_file.Close()

    codes[code] = true
    for k, v := range codes {
        if v {
            w_file.WriteString(k + " 1\n")
        } else {
            w_file.WriteString(k + " 0\n")
        }
    }
    return true
}


func Login(c *gin.Context) {
    var err error
    var result db.Result
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

    var user db.User

    result = DB.Where(
        "username=? and password=?",
        param.Username,
        param.Password,
    ).First(&user)

    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        c.JSON(http.StatusBadRequest, R.BadRequest(R.Json{
            "message": "Wrong username or password",
        }))
        return
    }

    var ret_token string
    ret_token, err = jwt.GenerateToken(user.UUID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, R.InternalServerError(nil))
        return
    }
    c.JSON(http.StatusOK, R.Success(R.Json{
        "token" : ret_token,
    }))
    return
}


func Register(c *gin.Context) {
    var err error

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

    // 检查激活码
    var codes, codeAvailable = checkActivationCode(param.Code)
    if !codeAvailable {
        c.JSON(
            http.StatusBadRequest,
            R.BadRequest(R.Json{"message":err.Error()}),
        )
        return
    }

    // // username判重
    // result := DB.Where("Username = ?", param.Username).Find(&db.User{})
    // if result.RowsAffected > 0 {
    //     c.JSON(
    //         http.StatusBadRequest,
    //         R.BadRequest(R.Json{"message": "Register failed"}),
    //     )
    //     return
    // }

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

    // 更新激活码
    var use_result = useActivationCode(codes, param.Code)
    if !use_result {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "Use activation code failed",
        })
        return
    }

    c.JSON(http.StatusOK, R.Success(R.Json{"token": ret_token}))
    return
}