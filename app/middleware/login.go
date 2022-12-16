package middleware

// 在e.Get中使用

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-svc-tpl/app/response"
)

// for Users

func CheckRegister(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// read
		name := c.FormValue("name")
		email := c.FormValue("email")
		pwd := c.FormValue("pwd")
		secQ := c.FormValue("secQ")
		secA := c.FormValue("secA")

		one := RegisterStruct{name, email, pwd, secQ, secA}
		valid := validator.New()
		invalid_err := valid.Struct(one)
		if invalid_err != nil {
			return response.SendResponse(c, 107, "invalid register info format")
		}
		return next(c)
	}
}

func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// read
		name := c.FormValue("email")
		pwd := c.FormValue("pwd")

		one := LoginStruct{name, pwd}
		valid := validator.New()
		invalid_err := valid.Struct(one)
		if invalid_err != nil {
			return response.SendResponse(c, 105, "invalid login info format")
		}
		return next(c)
	}
}

func CheckSecure(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// read
		name := c.FormValue("name")
		email := c.FormValue("email")
		secA := c.FormValue("secA")
		new_pwd := c.FormValue("newpwd")

		one := SecurityStruct{name, email, new_pwd, secA}
		valid := validator.New()
		invalid_err := valid.Struct(one)
		if invalid_err != nil {
			return response.SendResponse(c, 211, "invalid verification inputs")
		}
		return next(c)
	}
}

func CheckReset(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// read
		name := c.FormValue("name")
		new_pwd := c.FormValue("newpwd")
		old_pwd := c.FormValue("oldpwd")

		one := ResetStruct{name, new_pwd, old_pwd}
		valid := validator.New()
		invalid_err := valid.Struct(one)
		if invalid_err != nil {
			return response.SendResponse(c, 622, "old password found incorrect")
		}
		return next(c)
	}
}
