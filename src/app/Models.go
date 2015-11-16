package main
import (
	"time"
	"encoding/json"
)
type Student struct {
	Id        int            `gorm:"primary_key", sql:"AUTO_INCREMENT"` // set primary key
	Name      string 		`form:"name" binding:"required"`			// require if string '',null & int 0 & boolean false
	Address   string 		`form:"address" binding:"required"`			// form: "name" is the param when post data
	Class     string        `form:"class" binding:"required"`
	Age       int           `form:"age" binding:"required"`
	Gender    bool			`form:"gender"`
	CreatedAt time.Time      `sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time      `sql:"DEFAULT:current_timestamp"`
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