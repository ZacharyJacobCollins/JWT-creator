package main

import (
    "encoding/base64"
    "encoding/hex"
    "fmt"
    "log"
    "time"

    "github.com/dgrijalva/jwt-go"
)

func main() {
    apiKey := "key"
    apiSecret := "Mysecret"

    signingKey, err := base64.URLEncoding.DecodeString(apiSecret)
    if nil != err {
        log.Fatal(err)
    }
    token := jwt.New(jwt.SigningMethodHS256)
    token.Claims["aud"] = apiKey
    token.Claims["iat"] = time.Now().Unix()

    id := make([]byte, 16)
    if err != nil {
        log.Fatal(err)
    }
    token.Claims["jti"] = hex.EncodeToString(id)
    token.Claims["scopes"] = map[string]interface{}{
        "users": map[string]interface{}{
            "actions": []string{"create"},
        },
    }
    tokenString, err := token.SignedString(signingKey)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(tokenString)
}
