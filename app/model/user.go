package model

type User struct {
	BaseModel
	Nickname     string `json:"nickname" gorm:"column:nickname;type:varchar(255);"`
	Avatar       string `json:"avatar" gorm:"column:avatar;type:varchar(500);"`
	HashPassword string `json:"hash_password" gorm:"column:hash_password;type:varchar(64);"`
}
