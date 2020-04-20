package models

import "time"
/*
|-------------------------------------------------------
| Main result model
|-------------------------------------------------------
*/

type Result struct {
	ID        uint `json:"id" gorm:"primary_key"`
	UrlID      uint `json:"url_id" gorm:"NOT NULL"`
	StatusCode uint `json:"status_code" gorm:"NOT NULL"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}