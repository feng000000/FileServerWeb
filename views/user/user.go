package user

import (
    "net/http"
    "errors"
    "github.com/gin-gonic/gin"

    "FileServerWeb/db"
    L "FileServerWeb/widget/logger"
    R "FileServerWeb/widget/response"
)


var DB = db.DB


func ChangeUsernameHandler(c *gin.Context) {
    var result db.Result

    var UUID = c.GetString("UUID")

    type ChangeUsernameParams struct {
        NewUsername    string `json:"new_username" binding:"required"`
    }

    var param ChangeUsernameParams
    if c.ShouldBind(&param) != nil {
        c.JSON(http.StatusBadRequest, R.BadRequest(nil))
        return
    }

    var user = db.User{
        Username: param.NewUsername,
    }

    result = DB.Model(&db.User{}).Where("UUID = ?", UUID).Updates(user)
    if result.RowsAffected == 0{
        c.JSON(http.StatusBadRequest, R.BadRequest(R.Json{
            "message": "User dont exists",
        }))
        return
    } else if result.Error != nil {
        L.Logger.Error(result.Error.Error())
        c.JSON(http.StatusInternalServerError, R.DatabaseError(nil))
        return
    }

    c.JSON(http.StatusOK, R.Success(nil))
    return
}


func ChangePasswordHandler(c *gin.Context) {
    var result db.Result
    var UUID = c.GetString("UUID")

    type ChangePasswordParams struct {
        NewPassword     string `json:"new_password" binding:"required"`
        OldPassword     string `json:"old_password" binding:"required"`
    }

    var param ChangePasswordParams
    if c.ShouldBind(&param) != nil {
        c.JSON(http.StatusBadRequest, R.BadRequest(nil))
        return
    }

    var user = db.User{
        Password: param.NewPassword,
    }

    result = DB.Model(&db.User{}).Where("UUID=? and password=?", UUID, param.OldPassword).Updates(user)
    if errors.Is(result.Error, db.ErrRecordNotFound) {
        c.JSON(http.StatusBadRequest, R.BadRequest(R.Json{
            "message": "User dont exists",
        }))
        return
    } else if result.RowsAffected == 0{
        c.JSON(http.StatusBadRequest, R.BadRequest(R.Json{
            "message": "Wrong Password",
        }))
        return
    } else if result.Error != nil {
        L.Logger.Error(result.Error.Error())
        c.JSON(http.StatusInternalServerError, R.DatabaseError(nil))
        return
    }

    c.JSON(http.StatusOK, R.Success(nil))
    return
}
