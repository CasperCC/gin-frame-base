package service

import (
	"fmt"
	"gin-frame-base/app/model"
)

type userService struct {
	baseService
}

// GetUserById 根据用户id查询详情
func (s *userService) GetUserById(id uint) (model.User, error) {
	if s.db == nil {
		return model.User{}, fmt.Errorf("db connection failed")
	}
	var user model.User
	err := s.db.Where("id = ?", id).First(&user).Error
	return user, err
}
