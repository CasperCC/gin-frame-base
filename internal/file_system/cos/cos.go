package cos

import (
	"gin-frame-base/internal/global"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

func GetConnection() (cosClient *cos.Client) {
	config := global.Config.FileSystem

	u, _ := url.Parse("https://" + config.BucketName + ".cos." + config.Region + ".myqcloud.com")
	s, _ := url.Parse("https://cos." + config.Region + ".myqcloud.com")
	baseUrl := &cos.BaseURL{BucketURL: u, ServiceURL: s}
	cosClient = cos.NewClient(baseUrl, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.SecretId,
			SecretKey: config.SecretKey,
		},
	})

	return
}
