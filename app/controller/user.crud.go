package controller

// user crud
import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/app/middleware"
	"go-svc-tpl/app/response"
	"go-svc-tpl/model"
	"time"
)

// {host}/user/register (/:name/:email/:pwd/:secQ/:secA)

var status string

// {host}/user
func Users_Judge(c echo.Context) error {
	if status == "nil" {
		return response.SendResponse(c, 900, "didn't login", status)
	}
	return response.SendResponse(c, 901, "login", status)
}

// {host}/user/register
func Users_register(c echo.Context) error {

	data := new(model.RegisterInput)
	if err := c.Bind(data); err != nil {
		logrus.Error("Bind Failed")
		return response.SendResponse(c, 400, "Bind Failed")
	}
	current_time := time.Now()
	new_user := new(model.Users)
	(*new_user).Name = (*data).Name
	//new_user := model.Users{1, name, email, pwd, secQ, secA, current_time}

	err := model.DB.Debug().Find(&new_user).Error
	if err == nil {
		status = "nil"
		return response.SendResponse(c, 100, "User's name been used", status)
	}
	(*new_user).Id = (*data).Id
	(*new_user).Email = (*data).Email
	(*new_user).Pwd = (*data).Pwd
	(*new_user).SecQ = (*data).SecQ
	(*new_user).SecA = (*data).SecA
	(*new_user).LatestTime = current_time

	name := (*data).Name
	email := (*data).Email
	pwd := (*data).Pwd
	secQ := (*data).SecQ
	secA := (*data).SecA

	one := middleware.RegisterStruct{name, email, pwd, secQ, secA}
	valid := validator.New()
	invalid_err := valid.Struct(one)
	if invalid_err != nil {
		return response.SendResponse(c, 107, "invalid register info format", name, email, pwd, secQ, secA)
	}
	// validate

	err = model.DB.Debug().Create(&new_user).Error
	if err != nil {
		status = "nil"
		return response.SendResponse(c, 000, "User create failed", status)
	}
	status = new_user.Name
	return response.SendResponse(c, 101, "User creating success", status)

}

// {host}/user/login (/:email/:pwd)
// -> query
func User_login(c echo.Context) error {

	data := new(model.LoginInput)
	if err := c.Bind(data); err != nil {
		logrus.Error("Bind Failed")
		return response.SendResponse(c, 400, "Bind Failed")
	}
	return response.SendResponse(c, -100, "test", (*data).Email)

	a_User := new(model.Users)
	(*a_User).Email = (*data).Email
	(*a_User).Pwd = (*data).Pwd
	pwd := a_User.Pwd
	email := a_User.Email

	//fmt.Println(email)
	one := middleware.LoginStruct{email, pwd}
	valid := validator.New()
	invalid_err := valid.Struct(one)
	if invalid_err != nil {
		return response.SendResponse(c, 105, "invalid login info format", email, pwd)
	}

	err := model.DB.Debug().Find(&a_User).Error
	if err != nil {
		status = "nil"
		return response.SendResponse(c, 102, "incorrect email or password", status)
	}
	//if (*a_User).Pwd != pwd {
	//	status = "nil"
	//	return response.SendResponse(c, 102, "incorrect email or password", status)
	//}
	status = a_User.Name
	return response.SendResponse(c, 103, "login successfully", status)
}

// {host}/user/logout
func User_logout(c echo.Context) error {
	if status == "nil" {
		return response.SendResponse(c, 900, "didn' login", status)
	}
	status = "nil"
	return response.SendResponse(c, 900, "didn't login", status)
}

// {host}/user/security
func User_reset_pwd(c echo.Context) error {

	data := new(model.SecureResetPwdInput)
	if err := c.Bind(data); err != nil {
		logrus.Error("Bind Failed")
	}

	//va_user := new(middleware.SecurityStruct)
	name := (*data).Name
	email := (*data).Email
	secA := (*data).SecA
	new_pwd := (*data).Pwd_new

	one := middleware.SecurityStruct{name, email, new_pwd, secA}
	valid := validator.New()
	invalid_err := valid.Struct(one)
	if invalid_err != nil {
		return response.SendResponse(c, 211, "invalid verification inputs", name, email, new_pwd, secA)
	}

	a_user := new(model.Users)
	(*a_user).Name = (*data).Name
	(*a_user).Email = (*data).Email

	err := model.DB.Debug().Find(&a_user).Error
	if err != nil {
		return response.SendResponse(c, 221, "no user found or unpaired Name and Email")
	}

	if (*a_user).SecA != (data).SecA {
		return response.SendResponse(c, 200, "incorrect answer", status)
	}
	(*a_user).Pwd = (*data).Pwd_new
	return response.SendResponse(c, 201, "pwd reset success", status)
}

// {host}/user/info

func User_get(c echo.Context) error {
	name := status
	if status == "nil" {
		return response.SendResponse(c, 900, "didn't login", status)
	}
	a_user := model.Users{}
	a_user.Name = name
	err := model.DB.Debug().Find(&a_user).Error
	if err != nil {
		return response.SendResponse(c, 111, "no User (FATAL)", status)
	}
	email := a_user.Email
	return response.SendResponse(c, 112, "query succeeded", name, email)
}

// {host}/user/record/get
// IP 地址获取办法还有些疑惑
func User_login_get(c echo.Context) error {
	if status == "nil" {
		return response.SendResponse(c, 900, "didn't login", status)
	}
	name := status
	a_user := model.Users{}
	a_user.Name = name
	err := model.DB.Debug().Find(&a_user).Error
	if err != nil {
		return response.SendResponse(c, 111, "no User (FATAL)", status)
	}
	return response.SendResponse(c, 310, "getting information succeeded", a_user.LatestTime, a_user.Name, status)
}

// {host}/user/pwdreset
func User_pwd_reset(c echo.Context) error {
	if status == "nil" {
		return response.SendResponse(c, 900, "didn't login", status)
	}

	data := new(model.ResetPwdInput)
	if err := c.Bind(data); err != nil {
		logrus.Error("Bind Failed")
	}
	// 因为要用到middleware，我们还是要求输入name

	name := (*data).Name
	new_pwd := (*data).Pwdnew
	old_pwd := (*data).Pwdold

	one := middleware.ResetStruct{name, new_pwd, old_pwd}
	valid := validator.New()
	invalid_err := valid.Struct(one)
	if invalid_err != nil {
		return response.SendResponse(c, 622, "old password found incorrect", name)
	}

	a_user := new(model.Users)
	(*a_user).Email = (*data).Email
	(*a_user).Name = (*data).Name

	err := model.DB.Debug().Find(&a_user).Error
	if err != nil {
		return response.SendResponse(c, 599, "no user found")
	}
	if (*a_user).Pwd != (*data).Pwdold {
		return response.SendResponse(c, 600, "incorrect old password", status)
	}
	(*a_user).Pwd = (*data).Pwdnew
	return response.SendResponse(c, 601, "password reset succeeded", status)
}
