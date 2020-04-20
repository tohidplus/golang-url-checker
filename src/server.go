package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tohidplus/url_project/src/database"
	"github.com/tohidplus/url_project/src/exception"
	"github.com/tohidplus/url_project/src/helpers/url"
	"github.com/tohidplus/url_project/src/router"
	"log"
	"net/http"
)

func init() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3307)/url_tester?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		exception.LogPanic(err)
	}
	database.DB = db
	database.Migrate()
}
func main() {
	defer database.DB.Close()
	url.RunSchedule()
	log.Fatal(http.ListenAndServe(":8085", router.RegisterHttpRoutes()))
}
