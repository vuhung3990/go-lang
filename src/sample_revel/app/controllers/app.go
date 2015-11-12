package controllers

import (
	"github.com/revel/revel"
// please import mysql driver (if not have please install)
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// get constant from setting
	revel.INFO.Println(revel.Config.StringDefault("db.database", "ban anh trien"))

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
