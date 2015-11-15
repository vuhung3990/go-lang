package main
import (
	"time"
	"encoding/json"
)
type Student struct {
	Id                   int            `gorm:"primary_key", sql:"AUTO_INCREMENT"` // set primary key
	Name, Address, Class string
	Age                  int
	Gender               bool
	CreatedAt            time.Time      `sql:"DEFAULT:current_timestamp"`
	UpdatedAt            time.Time      `sql:"DEFAULT:current_timestamp"`
}

type JsonResponse struct {
	IsSuccess    bool
	Action, Data string


}

// set data json
func (j JsonResponse)setJsonResponseData(data interface{}) JsonResponse {
	json, _ := json.Marshal(data)
	j.Data = string(json)
	return j
}