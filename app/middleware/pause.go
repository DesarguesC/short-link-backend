package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Pauseurl(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 使用 echo.Context.Abort 函数终止重定向
		// 返回一个错误
		return c.String(http.StatusBadRequest, "Invalid request")
	}
}
