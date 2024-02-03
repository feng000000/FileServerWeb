// 用法
// var s string
// s, err = jwt.GenerateToken("feng")
// if err != nil {
//     fmt.Println("generate jwt failed, ", err)
//     return
// }
// fmt.Printf("token: %s\n", s)

// // 解析jwt
// var claim *jwt.Claims
// claim, err = jwt.ParseToken(s)
// if err != nil {
//     fmt.Println("parse jwt failed:", err)
//     return
// }
// fmt.Printf("claim: %+v\n", claim)
// fmt.Println("username: ", claim.Username)
package jwt


import (
    "errors"
    "strings"
    "time"

    "github.com/golang-jwt/jwt/v5"

    "FileServerWeb/config"
    "FileServerWeb/db"
	L "FileServerWeb/widget/logger"
)


var DB = db.DB


type Claims struct {
    UUID string `json:"uuid"`
    jwt.RegisteredClaims  // github.com/golang-jwt/jwt/ v5版本新加的方法
}


func GenerateToken(uuid string) (string, error) {
    var current = time.Now()
    var claim = Claims{
        uuid,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(current.Add(3 * time.Hour)), // 过期时间3小时
            IssuedAt:  jwt.NewNumericDate(current), // 签发时间
            NotBefore: jwt.NewNumericDate(current), // 生效时间
        },
    }

	// 通过 UUID 查询 SecretKey
	var user db.UserSecretKey

	// 使用 Where 条件进行查询
	result := DB.Where("uuid = ?", uuid).First(&user)
    if result.Error != nil {
        return "", result.Error
    }


    var t = jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
    // 必须传[]byte类型
    var s, err = t.SignedString([]byte(config.SECRET_KEY + user.SecretKey))
    L.Logger.Info(config.SECRET_KEY + user.SecretKey)


    return "Bearer " + s, err
}


func ParseToken(s string) (*Claims, error) {
    if s == "" {
        return nil, errors.New("Not login")
    }
    var res = strings.Split(s, " ")

    if res[0] != "Bearer" {
        return nil, errors.New("Wrong format")
    }

    // var t, err = jwt.ParseWithClaims(
    //     res[1],
    //     &Claims{},
    //     func(token *jwt.Token) (interface{}, error) {
    //         return []byte(config.SECRET_KEY), nil
    //     // func(token *jwt.Token) (interface{}, error) {
    //     //     return config.SECRET_KEY, nil
    // })


    // 解析 Token 字符串
    var t, err = jwt.ParseWithClaims(
        res[1],
        &Claims{},
        func(token *jwt.Token) (interface{}, error) {
            // 获取 Token 中的 UUID
            claims, ok := token.Claims.(*Claims)
            if !ok {
                return nil, errors.New("Failed to parse claims")
            }
            uuid := claims.UUID

            // 通过 UUID 查询 SecretKey
            var user db.UserSecretKey
            result := DB.Where("uuid = ?", uuid).First(&user)
            if result.Error != nil {
                return "", result.Error
            }

            // 返回用户密钥用于验证 Token 的签名
            return []byte(config.SECRET_KEY + user.SecretKey), nil
    })

    var claims, ok = t.Claims.(*Claims)
    if ok && t.Valid {
        return claims, nil
    } else {
        return nil, err
    }
}
