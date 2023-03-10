package model

type WordDefinition struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	WordName string `json:"word_name"`
	Meaning  string `json:"meaning"`
	Keyword  string `json:"keyword"`
}
