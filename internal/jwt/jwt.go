package jwt

import (
	"fmt"
	"gin-frame-base/internal/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User interface {
	GetUid() string
}

type CustomerClaims struct {
	jwt.StandardClaims
}

const TokenType = "bearer"

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func GenerateToken(user User) (tokenOutPut *TokenOutPut, token *jwt.Token, err error) {
	currentTimestamp := time.Now().Unix()
	token = jwt.NewWithClaims(
		jwt.SigningMethodES256,
		CustomerClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + global.Config.Jwt.Ttl,
				Id:        user.GetUid(),
				Issuer:    global.Config.App.Name,
				NotBefore: currentTimestamp - 1000,
			},
		},
	)
	fmt.Println(global.Config.Jwt.Secret)

	tokenStr, err := token.SignedString([]byte(global.Config.Jwt.Secret))

	tokenOutPut = &TokenOutPut{
		AccessToken: tokenStr,
		ExpiresIn:   int(currentTimestamp + global.Config.Jwt.Ttl),
		TokenType:   TokenType,
	}
	return
}
