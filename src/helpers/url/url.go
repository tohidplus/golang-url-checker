package url

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/carlescere/scheduler"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"url_project/src/app/models"
	"url_project/src/database"
	"url_project/src/exception"
)

func Call(url models.Url) *http.Response {
	client := http.Client{}
	body, _ := json.Marshal(url.Body)
	bodyReader := bytes.NewReader(body)
	request, _ := http.NewRequest(url.Method, url.Path, bodyReader)
	for key, value := range url.Headers {
		sValue, _ := value.(string)
		request.Header.Set(key, sValue)
	}
	res, _ := client.Do(request)
	fmt.Printf("The url %s with status %d\r\n", url.Path, res.StatusCode)
	return res
}

func RunSchedule() {
	scheduleMethods := scheduleMethods()
	var urls []models.Url
	database.DB.Find(&urls)
	i := 0
	for i < len(urls) {
		url := urls[i]
		job := getUrlJob(url)
		scheduleMethods[url.ScheduleType](url.ScheduleValue, job)
		i++
	}
	fmt.Println("Finished")
}

func getUrlJob(url models.Url) func() {
	return func() {
		go func(){
			res := Call(url)
			result := models.Result{
				UrlID:      url.ID,
				StatusCode: uint(res.StatusCode),
			}
			database.DB.Create(&result)
			stringCode := strconv.Itoa(res.StatusCode)
			if string(stringCode[0]) != "2" {
				url.FailedCount += 1
				database.DB.Save(&url)
				if url.FailedCount >= url.Threshold {
					err := database.DB.Transaction(func(tx *gorm.DB) error {
						url.FailedCount = 0
						notification := models.Notification{
							UserID:  url.UserID,
							UrlID:   url.ID,
							Message: "Status " + strconv.Itoa(res.StatusCode),
						}
						tx.Create(&notification)
						tx.Save(&url)
						return nil
					})
					exception.LogFatal(err)
				}
			}
		}()
	}
}

func scheduleMethods() map[string]func(every uint, function func()) {
	return map[string]func(every uint, function func()){
		"day": func(every uint, function func()) {
			_, _ = scheduler.Every(int(every)).Day().Run(function)
		},
		"hours": func(every uint, function func()) {
			_, _ = scheduler.Every(int(every)).Hours().Run(function)
		},
		"minutes": func(every uint, function func()) {
			_, _ = scheduler.Every(int(every)).Minutes().Run(function)
		},
		"seconds": func(every uint, function func()) {
			_, _ = scheduler.Every(int(every)).Seconds().Run(function)
		},
	}
}
