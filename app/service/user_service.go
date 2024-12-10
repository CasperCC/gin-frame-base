package service

import (
	"gin-frame-base/app/dao"
	"gin-frame-base/app/model"
	"gin-frame-base/app/request/user_request"
	"gin-frame-base/internal/jwt"
)

type UserService struct {
	userDao dao.UserDao
}

func NewUserService(userDao dao.UserDao) UserService {
	return UserService{userDao: userDao}
}

// GetUserById 根据用户id查询详情
func (s *UserService) GetUserById(id uint) (user model.User, err error) {
	user, err = s.userDao.GetByID(id)
	return
}

func (s *UserService) Login(params user_request.Login) (jwtToken *jwt.TokenOutPut, err error) {
	var user *model.User
	user, err = s.userDao.GetByCondition(0, params.Mobile)
	if user == nil || user.ID.ID == 0 {
		return nil, err
	}
	jwtToken, _, err = jwt.GenerateToken(user)
	return
}
