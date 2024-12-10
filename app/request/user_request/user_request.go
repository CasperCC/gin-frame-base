package user_request

type GetUserDetail struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

type Login struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Register struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required"`
	Avatar          string `form:"avatar" json:"avatar"`
	Mobile          string `form:"mobile" json:"mobile" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required"`
}
