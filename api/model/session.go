package model

type Session struct {
	ID     uint       `gorm:"primarykey" json:"id"`
	UserID uint       `json:"user_id"`
	User   User       `json:"user"`
	TagID  uint       `json:"tag_id"`
	Tag    SessionTag `gorm:"foreignKey:TagID" json:"tag"`
}
