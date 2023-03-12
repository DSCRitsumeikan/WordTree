package model

type WordNodeChildren struct {
	ID            uint     `gorm:"primarykey" json:"id"`
	CurrentNodeID uint     `json:"current_node_id"`
	CurrentNode   WordNode `json:"current_node"`
	ChildNodeID   uint     `json:"child_node_id"`
	ChildNode     WordNode `json:"child_node"`
}
