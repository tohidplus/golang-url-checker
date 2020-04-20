package models

import "time"

type Notification struct {
	ID        uint `json:"id" gorm:"primary_key"`
	UserID  uint   `json:"user_id" gorm:"NOT NULL"`
	UrlID   uint   `json:"url_id" gorm:"NOT NULL"`
	Message string `json:"message" gorm:"type:varchar(255); NOT NULL"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}