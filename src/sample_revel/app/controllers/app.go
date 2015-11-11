package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// get constant from setting
	revel.INFO.Println(revel.Config.StringDefault("db.database", "ban anh trien"))

	return c.Render()
}
