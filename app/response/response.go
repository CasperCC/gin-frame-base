package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	Json(c, CODE_SUCCESS, MSG_SUCCESS, data, http.StatusOK)
}

func Json(c *gin.Context, code int, msg string, data interface{}, httpStatus int) {
	response := Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	c.JSON(httpStatus, response)
}

func Error(c *gin.Context, code int, msg string, httpStatus ...int) {
	httpCode := http.StatusBadRequest
	if len(httpStatus) > 0 {
		httpCode = httpStatus[0]
	}
	Json(c, code, msg, nil, httpCode)
}
