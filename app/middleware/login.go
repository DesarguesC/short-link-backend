package middleware

// 在e.Get中使用

import (
	"github.com/labstack/echo/v4"
)

// for Users

//type RegisterInput struct {
//	Id    int    `gorm:"type:uint;primaryKey autoincrement=1000" form:"id" json:"id"`
//	Name  string `gorm:"type:varchar(20)" form:"name" json:"name"`
//	Email string `gorm:"type:varchar(50)" form:"email" json:"email"`
//	Pwd   string `gorm:"type:varchar(90)" form:"pwd" json:"pwd"`
//	SecQ  string `gorm:"type:varchar(100)" form:"secq" json:"secq"`
//	SecA  string `gorm:"type:varchar(100)" form:"seca" json:"seca"`
//}

func CheckRegister(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// read

		//name := c.QueryParam("name")
		//email := c.QueryParam("email")
		//pwd := c.QueryParam("pwd")
		//secQ := c.QueryParam("secQ")
		//secA := c.QueryParam("secA")

		//data := new(RegisterInput)
		//if err := c.Bind(data); err != nil {
		//	logrus.Error("Bind Failed")
		//}

		return next(c)
	}
}

func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// read
		//name := c.QueryParam("email")
		//pwd := c.QueryParam("pwd")
		//
		//one := LoginStruct{name, pwd}
		//valid := validator.New()
		//invalid_err := valid.Struct(one)
		//if invalid_err != nil {
		//	return response.SendResponse(c, 105, "invalid login info format", name)
		//}
		return next(c)
	}
}

//func Secure_SendQuestion(next echo.Context) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		name := c.QueryParam()("name")
//		a_user := new(model.Users)
//	}
//}

func CheckSecure(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// read
		//name := c.QueryParam("name")
		//email := c.QueryParam("email")
		//secA := c.QueryParam("secA")
		//new_pwd := c.QueryParam("newpwd")
		//
		//one := SecurityStruct{name, email, new_pwd, secA}
		//valid := validator.New()
		//invalid_err := valid.Struct(one)
		//if invalid_err != nil {
		//	return response.SendResponse(c, 211, "invalid verification inputs", name, email)
		//}
		return next(c)
	}
}

func CheckReset(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// read
		//name := c.QueryParam("name")
		//new_pwd := c.QueryParam("newpwd")
		//old_pwd := c.QueryParam("oldpwd")
		//
		//one := ResetStruct{name, new_pwd, old_pwd}
		//valid := validator.New()
		//invalid_err := valid.Struct(one)
		//if invalid_err != nil {
		//	return response.SendResponse(c, 622, "old password found incorrect", name)
		//}
		return next(c)
	}
}
