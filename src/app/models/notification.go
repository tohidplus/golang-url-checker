package models

import "github.com/jinzhu/gorm"

type Notification struct {
	gorm.Model
	UserID  uint   `json:"user_id" gorm:"NOT NULL"`
	UrlID   uint   `json:"url_id" gorm:"NOT NULL"`
	Message string `json:"message" gorm:"type:varchar(255); NOT NULL"`
}