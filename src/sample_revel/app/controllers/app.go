package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

//note: name is same with view/App/Hello.html
func (c App) Hello() revel.Result {
	return c.Render()
}
