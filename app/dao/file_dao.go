package dao

import "gin-frame-base/internal/global"

type FileDao struct {
	baseDao
}

func NewFileDao() FileDao {
	return FileDao{baseDao{global.Db}}
}
