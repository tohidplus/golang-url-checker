package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type JsonProperty map[string]interface{}

/*
|-------------------------------------------------------
| Main url  model
|-------------------------------------------------------
*/

type Url struct {
	ID            uint           `json:"id" gorm:"primary_key"`
	UserID        uint           `json:"user_id" gorm:"NOT NULL"`
	Path          string         `json:"path" gorm:"type:varchar(255); NOT NULL"`
	Method        string         `json:"method" gorm:"type:enum('GET','POST','PATCH','PUT','DELETE')"`
	Headers       JsonProperty   `json:"headers" gorm:"type:json; NOT NULL"`
	Body          JsonProperty   `json:"body" gorm:"type:json; NOT NULL"`
	ScheduleType  string         `json:"schedule_type"`
	ScheduleValue uint            `json:"schedule_value"`
	Threshold     uint           `json:"threshold" gorm:"NOT NULL"`
	FailedCount   uint           `json:"failed_count" gorm:"DEFAULT:0; NOT NULL"`
	Notifications []Notification `json:"notifications" gorm:"foreignKey:UrlID"`
	Results       []Result       `json:"results" gorm:"foreignKey:UrlID"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     *time.Time     `json:"deleted_at" sql:"index"`
}

/*
|-------------------------------------------------------
| Mutator
|-------------------------------------------------------
*/
func (jp JsonProperty) Value() (driver.Value, error) {
	j, err := json.Marshal(jp)
	return j, err
}

func (jp *JsonProperty) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	if err := json.Unmarshal(source, &i); err != nil {
		return err
	}

	*jp, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}
	return nil
}
