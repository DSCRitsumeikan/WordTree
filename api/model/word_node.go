package model

import "time"

type WordNode struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	UUID             string         `json:"uuid"`
	UserID           uint           `json:"user_id"`
	User             User           `json:"user"`
	ParentNodeID     uint           `json:"parent_node_id"`
	ParentNode       *WordNode      `json:"parent_node"`
	LeftNodeID       uint           `json:"left_node_id"`
	LeftNode         *WordNode      `json:"left_node"`
	RightNodeID      uint           `json:"right_node_id"`
	RightNode        *WordNode      `json:"right_node"`
	WordDefinitionID uint           `json:"word_definition_id"`
	WordDefinition   WordDefinition `json:"word_definition"`
	CreatedAt        time.Time      `gorm:"autoCreateTime:milli" json:"created_at"`
}
