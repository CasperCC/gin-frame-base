package model

type File struct {
	ID
	UserId         uint   `json:"user_id" gorm:"column:user_id;type:int(11)"`
	OriginFilename string `json:"origin_filename" gorm:"column:origin_filename;type:varchar(255);"`
	FileKey        string `json:"file_key" gorm:"column:file_key;type:varchar(255);"`
	Size           int64  `json:"size" gorm:"column:size;type:bigint(20);"`
	MimeType       string `json:"mime_type" gorm:"column:mime_type;type:varchar(255);"`
	Status         string `json:"status" gorm:"column:status;type:tinyint(3)"`
	FileHash       string `json:"file_hash" gorm:"column:file_hash;type:varchar(64);"`
	Timestamps
	SoftDelete
}
