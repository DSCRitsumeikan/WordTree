package model

// 使うケースがわからない
type CurrentNodeRelation struct {
	ID            uint     `gorm:"primarykey" json:"id"`
	CurrentNodeID uint     `json:"current_node_id"`
	CurrentNode   WordNode `json:"current_node"`
	LeftNodeID    uint     `json:"left_node_id"`
	LeftNode      WordNode `json:"left_node"`
	RightNodeID   uint     `json:"right_node_id"`
	RightNode     WordNode `json:"right_node"`
}
