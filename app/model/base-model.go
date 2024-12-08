package model

type BaseModel struct {
	ID        uint  `json:"id" gorm:"primary_key" json:"id"`
	CreatedAt int32 `json:"created_at" gorm:"column:created_at;type:int(11);"`
	UpdatedAt int32 `json:"updated_at" gorm:"column:updated_at;type:int(11);"`
	DeletedAt int32 `json:"deleted_at" gorm:"column:deleted_at;type:int(11);"`
}
