package middleware

import (
	"gin-frame-base/app/dao"
	"gin-frame-base/app/response"
	"gin-frame-base/internal/global"
	jwtUtil "gin-frame-base/internal/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
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

	// Token 续签
	if claims.ExpiresAt.Unix()-time.Now().Unix() < global.Config.Jwt.RefreshGracePeriod {
		userId, err := strconv.ParseUint(claims.ID, 10, 32)
		if err != nil {
			response.Error(c, response.CODE_SYSTEM_ERROR, response.MSG_SYSTEM_ERROR)
			c.Abort()
			return
		}
		userDao := dao.NewUserDao()
		user, err := userDao.GetByID(uint(userId))
		newToken, _, err := jwtUtil.GenerateToken(&user)
		if err != nil || newToken == nil {
			response.Error(c, response.CODE_DB_ERROR, response.MSG_DB_ERROR)
			c.Abort()
			return
		}
		c.Header("new_access_token", newToken.AccessToken)
		c.Header("new_expires_in", strconv.Itoa(newToken.ExpiresIn))
	}

	c.Set("jwt_token", token)
	c.Set("user_id", claims.ID)

	c.Next()
}
