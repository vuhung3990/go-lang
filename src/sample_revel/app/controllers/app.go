package controllers

import (
	"github.com/revel/revel"
	"time"
)

type Application struct {
	GormController
}

func (c Application) Index() revel.Result {
	// get constant from setting
	revel.INFO.Println(revel.Config.StringDefault("db.database", "ban anh trien"))

	//	c.db.NewRecord(model.User{Name:"hung vu", Address:"thai hoa", Age:25})
	return c.Render()
}

type Student struct {
	// model in gorm start with upper case
	Name      string
	Class     string
	Score     int
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
