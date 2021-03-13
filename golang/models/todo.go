package models

import "time"

type Todo struct {
	UUID      string    `gorm:"type:varchar(36); primary_key; not null"       json:"uuid"`
	Title     string    `gorm:"type:varchar(200);not null" json:"title"`
	State     bool      `gorm:"not null"                   json:"state"`
	Deadline  time.Time `gorm:"not null" json:"deadline"`
	CreatedAt time.Time `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

type CreateTodoInput struct {
	Title    string    `json:"title" binding: "required"`
	Deadline time.Time `json:"deadline" binding: "required"`
}
