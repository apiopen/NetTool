package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Favicon("./assets/favicon.ico")
	app.StaticWeb("/", "./ui")
	app.RegisterView(iris.HTML("./ui", ".html").Reload(true))
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	app.Done(func(ctx iris.Context) {})
	app.Run(iris.Addr(":80"), iris.WithCharset("UTF-8"))
}
