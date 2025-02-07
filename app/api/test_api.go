package api

import (
	"fmt"
	"gin-frame-base/app/response"
	"gin-frame-base/internal/global"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type TestApi struct{}

func (t *TestApi) Buckets(c *gin.Context) {
	s, _, err := global.Cos.Service.Get(c)
	if err != nil {
		response.Error(c, response.CODE_SYSTEM_ERROR, err.Error())
		return
	}
	response.Success(c, s.Buckets)
}

func (t *TestApi) Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		response.Error(c, response.CODE_SYSTEM_ERROR, "文件获取失败"+err.Error())
		return
	}

	// 2. 打开文件
	file, err := fileHeader.Open()
	if err != nil {
		response.Error(c, response.CODE_SYSTEM_ERROR, "文件打开失败"+err.Error())
		return
	}
	defer file.Close()

	// 3. 上传到 COS
	objectKey := "test_uploads/" + fileHeader.Filename // 存储路径
	_, err = global.Cos.Object.Put(c, objectKey, file, nil)
	if err != nil {
		response.Error(c, response.CODE_SYSTEM_ERROR, err.Error())
		return
	}
	response.Success(c, nil)
}

func (t *TestApi) FileList(c *gin.Context) {
	opt := &cos.BucketGetOptions{
		Prefix:    global.Config.FileSystem.Prefix,
		Delimiter: "/",
		MaxKeys:   1000,
	}
	isTruncated := true
	var marker string
	for isTruncated {
		opt.Marker = marker
		v, _, err := global.Cos.Bucket.Get(c, opt)
		if err != nil {
			fmt.Println(err)
			break
		}
		response.Success(c, v)
		return
		//for _, content := range v.Contents {
		//	fmt.Printf("Object: %v\n", content.Key)
		//}
		//// common prefix 表示表示被 delimiter 截断的路径, 如 delimter 设置为/, common prefix 则表示所有子目录的路径
		//for _, commonPrefix := range v.CommonPrefixes {
		//	fmt.Printf("CommonPrefixes: %v\n", commonPrefix)
		//}
	}
}
