package database

import (
	"github.com/jinzhu/gorm"
	"url_project/src/app/models"
)

var DB *gorm.DB

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Url{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	DB.AutoMigrate(&models.Result{}).AddForeignKey("url_id", "urls(id)", "cascade", "cascade")
	DB.AutoMigrate(models.Notification{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").
		AddForeignKey("url_id", "urls(id)", "cascade", "cascade")
}
