package user_request

type GetUserDetail struct {
	ID uint `form:"id" json:"id" binding:"required"`
}
