package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	text := "Hello World!"
	return c.Render(text)
}

func (c App) Hoge() revel.Result {
	type Result struct {
		Message string `json: "message"`
	}

	result := Result{
		Message: "hogehoge",
	}

	return c.RenderJSON(result)
}
