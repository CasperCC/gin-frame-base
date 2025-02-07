package service

import "gin-frame-base/app/dao"

type FileService struct {
	fileDao dao.FileDao
}

func NewFileService(fileDao dao.FileDao) FileService {
	return FileService{fileDao: fileDao}
}
