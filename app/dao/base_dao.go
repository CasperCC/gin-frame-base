package dao

import "gorm.io/gorm"

type baseDao struct {
	db *gorm.DB
}
