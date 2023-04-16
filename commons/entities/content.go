package entities

import "time"

type (
	Content struct {
		Id        string    `json:"id" gorm:"primaryKey"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
		UpdateAt  time.Time `json:"update_at"`
	}
)
