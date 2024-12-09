package service

import "gorm.io/gorm"

type baseService struct {
	db *gorm.DB
}
