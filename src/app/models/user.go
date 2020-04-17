package models

import (
	"database/sql/driver"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Password string
/*
|-------------------------------------------------------
| Main user model
|-------------------------------------------------------
*/
type User struct {
	gorm.Model
	Name          string         `json:"name" gorm:"type:varchar(255); NOT NULL"`
	Email         string         `json:"email" gorm:"type:varchar(255); NOT NULL; unique_index"`
	Password      Password       `gorm:"type:varchar(255); NOT NULL"`
	Urls          []Url          `json:"urls" gorm:"foreignKey:UserID"`
	Notifications []Notification `json:"notifications" gorm:"foreignKey:UserID"`
}

func (pass Password) Value() (driver.Value, error) {
	bPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return nil, errors.New("Error while hashing the password.")
	}
	return string(bPass), nil
}

/*
|-------------------------------------------------------
| A model to response
|-------------------------------------------------------
*/

type ResponseUser struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (responseUser ResponseUser) TableName() string {
	return "users"
}
