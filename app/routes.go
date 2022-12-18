package app

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-svc-tpl/app/controller"
	"go-svc-tpl/app/middleware"
)

func ping(c echo.Context) error {
	return c.String(200, "pong!")
}

func addRoutes() {
	api := e.Group("api") // /api
	api.GET("/doc/*", echoSwagger.WrapHandler)
	api.GET("/ping", ping)
	api.POST("/url/create", controller.CreateUrl)
	api.POST("/url/query", controller.QueryUrl)
	api.POST("/url/update", controller.UpdateUrl)
	api.POST("/url/Delete", controller.DelUrl)
	api.POST("/url/Pause", controller.PauseUrl)
	api.POST("/url/Continue", controller.ContinueUrl)
	api.POST("/short/:hash", controller.Visit) //动态参数路由 /:hash

	api.POST("/user/register", controller.Users_register, middleware.CheckRegister)
	api.POST("/user/login", controller.User_login, middleware.CheckLogin)
	api.POST("/user/logout", controller.User_logout)
	api.POST("/user/security", controller.User_reset_pwd, middleware.CheckSecure)
	api.POST("/user/info", controller.User_get)
	api.POST("/user/record/get", controller.User_login_get)
	api.POST("/user/pwdreset", controller.User_pwd_reset, middleware.CheckReset)

}
