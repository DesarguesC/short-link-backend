package app

import (
	"go-svc-tpl/model"
	"short-link-backend/model"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ping(c echo.Context) error {
	return c.String(200, "pong!")
}

func addRoutes() {
	api := e.Group("api")
	api.GET("/doc/*", echoSwagger.WrapHandler)
	api.GET("/ping", ping)
	api.POST("/url/create", model.CreateUrl)
	api.POST("/url/query", model.QueryUrl)
	api.POST("/url/update", model.UpdateUrl)
	api.POST("/url/Delete", model.DelUrl)
}
