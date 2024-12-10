package middleware

import (
	"gin-frame-base/app/response"
	"gin-frame-base/internal/global"
	jwtUtil "gin-frame-base/internal/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuth(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	if tokenStr == "" || len(tokenStr) <= len(jwtUtil.TokenType)+1 {
		response.Error(c, response.CODE_UNAUTHORIZED, response.MSG_UNAUTHORIZED)
		c.Abort()
		return
	}
	tokenStr = tokenStr[len(jwtUtil.TokenType)+1:]

	// Token 解析校验
	token, err := jwt.ParseWithClaims(tokenStr, &jwtUtil.CustomerClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Secret), nil
	})
	if err != nil || token == nil || !token.Valid {
		response.Error(c, response.CODE_UNAUTHORIZED, response.MSG_UNAUTHORIZED)
		c.Abort()
		return
	}

	claims := token.Claims.(*jwtUtil.CustomerClaims)
	// Token 发布者校验
	if claims.Issuer != global.Config.App.Name {
		response.Error(c, response.CODE_UNAUTHORIZED, response.MSG_UNAUTHORIZED)
		c.Abort()
		return
	}

	c.Set("jwt_token", token)
	c.Set("user_id", claims.ID)

	c.Next()
}
