package models

import "github.com/jinzhu/gorm"

type Result struct {
	gorm.Model
	UrlID      uint `json:"url_id" gorm:"NOT NULL"`
	StatusCode uint `json:"status_code" gorm:"NOT NULL"`
}