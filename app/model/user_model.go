package model

import "strconv"

type User struct {
	ID
	Nickname     string `json:"nickname" gorm:"column:nickname;type:varchar(255);"`
	Avatar       string `json:"avatar" gorm:"column:avatar;type:varchar(500);"`
	Mobile       string `json:"mobile" gorm:"column:mobile;type:varchar(255);"`
	HashPassword string `json:"hash_password" gorm:"column:hash_password;type:varchar(64);"`
	Timestamps
	SoftDelete
}

func (u *User) GetUid() string {
	return strconv.Itoa(int(u.ID.ID))
}
