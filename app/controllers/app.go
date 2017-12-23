package controllers

import (
	"fmt"                              // 新規追加
	_ "github.com/go-sql-driver/mysql" // 新規追加
	"github.com/jinzhu/gorm"           // 新規追加
	"github.com/revel/revel"
	"strings" // 新規追加
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	connectionString := getConnectionString()

	// DB接続
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}
	db.DB()

	// ユーザー情報の作成
	user := &User{Name: "go taro", Age: 26}

	db.Create(user)

	ret := db.First(&user)

	return c.RenderJSON(ret)
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

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	host := getParamString("db.host", "")
	port := getParamString("db.port", "3306")
	user := getParamString("db.user", "")
	pass := getParamString("db.password", "")
	dbname := getParamString("db.name", "auction")
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("dbargs", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", user, pass, protocol, host, port, dbname, dbargs)
}
