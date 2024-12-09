package api

import (
	"gin-frame-base/app/request/user_request"
	"gin-frame-base/app/response"
	"gin-frame-base/app/service"
	"github.com/gin-gonic/gin"
)

type userApi struct{}

// GetUserDetail 获取用户详情
func (*userApi) GetUserDetail(c *gin.Context) {
	//user := model.User{BaseModel: model.BaseModel{ID: 1}}
	var request user_request.GetUserDetail
	if err := c.ShouldBind(&request); err != nil {
		response.Error(c, response.CODE_MISSING_PARAMS, response.MSG_MISSING_PARAMS)
		return
	}
	detail, err := service.UserService.GetUserById(request.ID)
	if nil != err {
		response.Error(c, response.CODE_DB_ERROR, err.Error())
		return
	}
	response.Success(c, detail)
	return
}
