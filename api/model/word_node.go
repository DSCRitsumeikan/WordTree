package model

import "time"

type WordNode struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	UUID             string         `json:"uuid"`
	UserID           uint           `json:"user_id"`
	User             User           `json:"user"`
	WordDefinitionID uint           `json:"word_definition_id"`
	WordDefinition   WordDefinition `json:"word_definition"`
	CreatedAt        time.Time      `gorm:"autoCreateTime:milli" json:"created_at"`
}
