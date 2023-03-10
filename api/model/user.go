package model

import "time"

type User struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	LastActiveAt time.Time `json:"last_active_at"`
}
