package jwt

import (
	"gin-frame-base/app/model"
	"gin-frame-base/internal/global"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type User interface {
	GetUser() *model.User
}

type CustomerClaims struct {
	jwt.RegisteredClaims
}

const TokenType = "bearer"

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func GenerateToken(user User) (tokenOutPut *TokenOutPut, token *jwt.Token, err error) {
	currentTimestamp := time.Now()
	claims := CustomerClaims{
		jwt.RegisteredClaims{
			Issuer:    global.Config.App.Name,
			Subject:   user.GetUser().Nickname,
			ExpiresAt: jwt.NewNumericDate(currentTimestamp.Add(time.Duration(global.Config.Jwt.Ttl) * time.Second)),
			NotBefore: jwt.NewNumericDate(currentTimestamp.Add(-1000 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(currentTimestamp.Add(-1000 * time.Second)),
			ID:        strconv.Itoa(int(user.GetUser().ID.ID)),
		},
	}
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenStr, err := token.SignedString([]byte(global.Config.Jwt.Secret))

	tokenOutPut = &TokenOutPut{
		AccessToken: tokenStr,
		ExpiresIn:   int(currentTimestamp.Unix() + global.Config.Jwt.Ttl),
		TokenType:   TokenType,
	}
	return
}
