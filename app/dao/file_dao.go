package dao

import (
	"gin-frame-base/app/model"
	"gin-frame-base/internal/global"
)

type FileDao struct {
	baseDao
}

func NewFileDao() FileDao {
	return FileDao{baseDao{global.Db}}
}

func (d *FileDao) GetById(id uint) (file *model.File, err error) {
	err = d.db.Where("id = ?", id).First(file).Error
	return file, err
}

func (d *FileDao) Create(file *model.File) error {
	return d.db.Create(file).Error
}

func (d *FileDao) Delete(file *model.File) error {
	return d.db.Where("fileKey = ?", file.FileKey).Delete(&model.File{}).Error
}
