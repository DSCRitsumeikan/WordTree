package model

import "time"

type CacheImage struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Keyword   string    `json:"keyword"`
	ImageHash string    `json:"image_hash"`
	Expire    time.Time `json:"expire"`
}
