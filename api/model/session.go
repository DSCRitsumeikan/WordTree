package model

import "time"

type Session struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user"`
}
