package util

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

// 接口身份校验jwt工具包

var jwtSecret []byte

type Claims struct {
    Username string `json:"username"`
    Password string `json:"password"`
    jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
    nowTime := time.Now()
    // 过期时间
    expireTime := nowTime.Add(30 * time.Minute)

    claims := Claims{
        EncodeMD5(username),
        EncodeMD5(password),
        jwt.StandardClaims{
            ExpiresAt: expireTime.Unix(),
            Issuer:    "gin-blog",
        },
    }

    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := tokenClaims.SignedString(jwtSecret)

    return token, err
}

func ParseToken(token string) (*Claims, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }
    return nil, err
}
