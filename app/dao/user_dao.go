package dao

import (
	"gin-frame-base/app/model"
	"gin-frame-base/internal/global"
)

type UserDao struct {
	baseDao
}

func NewUserDao() UserDao {
	return UserDao{baseDao{global.Db}}
}

func (u *UserDao) GetByID(id uint) (user model.User, err error) {
	err = u.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (u *UserDao) GetByCondition(id uint, mobile string) (*model.User, error) {
	var user model.User
	query := u.db.Model(&user)

	if id > 0 {
		query.Where("id = ?", id)
	}

	if mobile != "" {
		query.Where("mobile = ?", mobile)
	}

	err := query.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
