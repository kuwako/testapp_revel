package controllers

import (
	"fmt"                            // 新規追加
	"github.com/go-sql-driver/mysql" // 新規追加
	"github.com/jinzhu/gorm"         // 新規追加
	"github.com/revel/revel"
	"strings" // 新規追加
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

func (c App) Fuga(number int) revel.Result {
	type Result struct {
		Message string `json:"message"`
		Number  int    `json:"number"`
	}

	number = number * 10

	result := Result{
		Message: "url/:引数名とする引数を取れるようになります",
		Number:  number,
	}

	return c.RenderJSON(result)
}

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}
