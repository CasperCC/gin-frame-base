package api

import (
	"gin-frame-base/app/request/user_request"
	"gin-frame-base/app/response"
	"gin-frame-base/app/service"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	userService service.UserService
}

func NewUserApi(userService service.UserService) *UserApi {
	return &UserApi{userService}
}

// GetUserDetail
// 获取用户详情
//
//	@receiver a
//	@param c
func (a *UserApi) GetUserDetail(c *gin.Context) {
	var request user_request.GetUserDetail
	if err := c.ShouldBind(&request); err != nil {
		response.Error(c, response.CODE_MISSING_PARAMS, response.MSG_MISSING_PARAMS)
		return
	}
	detail, err := a.userService.GetUserById(request.ID)
	if nil != err {
		response.Error(c, response.CODE_DB_ERROR, err.Error())
		return
	}
	response.Success(c, detail)
	return
}

// Login
// 登录接口
//
//	@receiver a
//	@param c
func (a *UserApi) Login(c *gin.Context) {
	var request user_request.Login
	if err := c.ShouldBind(&request); err != nil {
		response.Error(c, response.CODE_MISSING_PARAMS, response.MSG_MISSING_PARAMS)
		return
	}
	jwtToken, err := a.userService.Login(request)
	if err != nil || jwtToken == nil {
		response.Error(c, response.CODE_LOGIN_FAILED, response.MSG_LOGIN_FAILED)
		return
	}
	response.Success(c, jwtToken)
}
