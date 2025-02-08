package api

import (
	"gin-frame-base/app/response"
	"gin-frame-base/app/service"
	"github.com/gin-gonic/gin"
)

type FileApi struct {
	fileService service.FileService
}

func NewFileApi(fileService service.FileService) *FileApi { return &FileApi{fileService} }

// Upload
// 上传文件接口
//
//	@receiver a
//	@param c
func (a *FileApi) Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		response.Error(c, response.CODE_SYSTEM_ERROR, "文件上传失败")
		return
	}
	_, err = a.fileService.UploadFile(c, fileHeader)
	if err != nil {
		response.Error(c, response.CODE_UPLOAD_FAILED, err.Error())
		return
	}
	response.Success(c, nil)
}
