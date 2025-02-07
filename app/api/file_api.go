package api

import "gin-frame-base/app/service"

type FileApi struct {
	fileService service.FileService
}

func NewFileApi(fileService service.FileService) *FileApi { return &FileApi{fileService} }
