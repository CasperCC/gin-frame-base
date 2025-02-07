package model

type ID struct {
	ID uint `json:"id" gorm:"primary_key" json:"id"`
}

type Timestamps struct {
	CreatedAt int32 `json:"created_at" gorm:"column:created_at;type:int(11);"`
	UpdatedAt int32 `json:"updated_at" gorm:"column:updated_at;type:int(11);"`
}

type SoftDelete struct {
	DeletedAt *int32 `json:"deleted_at" gorm:"column:deleted_at;type:int(11);"`
	DelId     int32  `json:"del_id" gorm:"column:del_id;type:int(11);"`
}
