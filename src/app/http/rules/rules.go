package rules

import (
	"encoding/json"
	"github.com/tohidplus/url_project/src/database"
	"regexp"
	"strconv"
	"strings"
)

func CheckEmail(email interface{}) bool {
	sEmail, ok := email.(string)
	if !ok {
		return false
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(sEmail)
}

func CheckUniqueness(table string, column string, value interface{}) bool {
	var count int
	database.DB.Table(table).Where(column+" = ?", value).Count(&count)
	return count == 0
}

func CheckUrl(url interface{}) bool {
	sUrl, ok := url.(string)
	if !ok {
		return false
	}
	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
	return re.MatchString(sUrl)
}

func CheckInArrayString(value string, array []string) bool {
	for _, v := range array {
		if strings.ToLower(value) == strings.ToLower(v) {
			return true
		}
	}
	return false
}

func CheckIsAnInteger(value interface{}) bool {
	sValue, ok := value.(string)
	if !ok {
		return false
	}
	_, cok := strconv.ParseInt(sValue, 10, 64)
	return cok == nil
}

func CheckJsonString(value interface{}) bool {
	sValue, ok := value.(string)
	if !ok {
		return false
	}
	var mValue map[string]json.RawMessage
	err := json.Unmarshal([]byte(sValue), &mValue)
	return err==nil
}