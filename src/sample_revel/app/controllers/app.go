package controllers

import (
	"github.com/revel/revel"
	"bytes"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

//note: name is same with view/App/Hello.html
func (c App) Hello() revel.Result {
	return c.RenderText("hell")
}

// post data sample
// note: every controller will start with Upper-Case
func (c App) CreateAccount(email, password, name string, age int) revel.Result {
	// validations
	c.Validation.Email(email).Message("email not valid")
	c.Validation.MinSize(password, 6).Message("password too short")
	c.Validation.Required(name).Message("name is required")
	c.Validation.Min(age, 3).Message("age too young, are you kidding me?")

	if (c.Validation.HasErrors()) {
		// if you want efficiently concatenate strings in Go, use buffer
		// you can use a+= "string...."

		var r bytes.Buffer
		for _, v := range c.Validation.Errors {
			r.WriteString(v.Message + "<br/>")
		}
		return c.RenderText(r.String())
	}

	return c.RenderText("success")
}
