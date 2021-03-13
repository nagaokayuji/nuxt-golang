package models

type Todo struct {
	ID    int    `gorm:"primary_key;not null"       json:"id"`
	Name  string `gorm:"type:varchar(200);not null" json:"name"`
	Memo  string `gorm:"type:varchar(400)"          json:"memo"`
	State int    `gorm:"not null"                   json:"state"`
}
