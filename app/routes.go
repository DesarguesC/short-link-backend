package app

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-svc-tpl/app/controller"
)

func ping(c echo.Context) error {
	return c.String(200, "pong!")
}

func addRoutes() {
	api := e.Group("api")
	api.GET("/doc/*", echoSwagger.WrapHandler)
	api.GET("/ping", ping)
	api.POST("/url/create", controller.CreateUrl)
	api.POST("/url/query", controller.QueryUrl)
	api.POST("/url/update", controller.UpdateUrl)
	api.POST("/url/Delete", controller.DelUrl)
}
