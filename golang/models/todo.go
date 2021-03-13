package models

import "time"

type Todo struct {
	UUID      string    `gorm:"type:varchar(36); primary_key; not null"       json:"uuid"`
	Title     string    `gorm:"type:varchar(200);not null" json:"title"`
	State     bool      `gorm:"not null"                   json:"state"`
	CreatedAt time.Time `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
