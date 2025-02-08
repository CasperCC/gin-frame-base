package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gin-frame-base/app/dao"
	"gin-frame-base/app/model"
	"gin-frame-base/internal/global"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"mime"
	"mime/multipart"
	"path/filepath"
)

type FileService struct {
	fileDao dao.FileDao
}

func NewFileService(fileDao dao.FileDao) FileService {
	return FileService{fileDao: fileDao}
}

// UploadFile
// 上传文件
//
//	@receiver s
//	@param c
//	@param fileHeader
//	@return err
func (s *FileService) UploadFile(c *gin.Context, fileHeader *multipart.FileHeader) (fileModel *model.File, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return
	}
	defer file.Close()

	// 获取文件大小（字节）
	fileSize := fileHeader.Size

	// 创建MD5哈希计算器
	hasher := md5.New()
	tee := io.TeeReader(file, hasher)

	// 生成唯一文件名
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return
	}

	originName := filepath.Base(fileHeader.Filename)
	ext := filepath.Ext(originName)
	baseName := originName[:len(originName)-len(ext)]

	userId, _ := c.Get("user_id")
	userCode, _ := c.Get("user_code")
	fileKey := fmt.Sprintf("%s%s/%s_%s%s", global.Config.FileSystem.Prefix, userCode.(string), baseName, uuidObj.String(), ext)

	_, err = global.Cos.Object.Put(c, fileKey, tee, nil)
	if err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			if fileKey != "" {
				_ = s.DeleteFile(c, fileKey)
			}
			// 继续抛出 panic
			panic(r)
		} else if err != nil && fileKey != "" {
			_ = s.DeleteFile(c, fileKey)
		}
	}()

	// 计算最终MD5值
	md5Hash := hex.EncodeToString(hasher.Sum(nil))
	mimeType := mime.TypeByExtension(ext)

	// 保存数据库
	fileModel = &model.File{
		UserId:         userId.(uint),
		OriginFilename: originName,
		FileKey:        fileKey,
		Size:           fileSize,
		MimeType:       mimeType,
		FileHash:       md5Hash,
	}
	err = s.fileDao.Create(fileModel)
	if err != nil {
		return
	}

	fileKey = ""
	return
}

// DeleteFile
// 删除文件
//
//	@receiver s
//	@param c
//	@param fileKey
//	@return err
func (s *FileService) DeleteFile(c *gin.Context, fileKey string) (err error) {
	isExist, err := global.Cos.Object.IsExist(c, fileKey)
	fmt.Println("isExist", isExist)
	if isExist && err == nil {
		_, err = global.Cos.Object.Delete(c, fileKey)
		if err != nil {
			return
		}
	}

	file := &model.File{FileKey: fileKey}
	err = s.fileDao.Delete(file)
	if err != nil {
		return
	}
	return
}
