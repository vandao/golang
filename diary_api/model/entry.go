package model

import (
	"diary_api/database"
)

type Entry struct {
	BaseModel
	UserId  uint   `json:"user_id"`
	Content string `gorm:"type:text;" json:"content"`
}

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}
