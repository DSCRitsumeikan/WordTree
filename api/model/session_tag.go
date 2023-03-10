package model

type SessionTag struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
	User   User   `json:"user"`
}
