package model

type Session struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
	User   User   `json:"user"`
}
