package controller

// user crud
import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/app/response"
	"go-svc-tpl/model"
	"time"
)

// {host}/user/register/:name/:email/:pwd/:secQ/:secA

var status string

// {host}/user
func Users_Judge(c echo.Context) error {
	if status == "nil" {
		return response.SendResponse(c, 900, "didn't login", status)
	}
	return response.SendResponse(c, 901, "login", status)
}

func Users_register(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	pwd := c.FormValue("pwd")
	secQ := c.FormValue("secQ")
	secA := c.FormValue("secA")
	current_time := time.Now()

	new_user := model.Users{}
	new_user.Name = name
	new_user.Email = email
	new_user.Pwd = pwd
	new_user.SecQ = secQ
	new_user.SecA = secA
	new_user.LatestTime = current_time

	//new_user := model.Users{1, name, email, pwd, secQ, secA, current_time}

	err := model.DB.Debug().Create(&new_user).Error
	if err != nil {
		status = "nil"
		return response.SendResponse(c, 100, "User creating error", status)
	}
	status = name
	return response.SendResponse(c, 101, "User creating seccess", status)
}

// {host}/url/login/:email/:pwd
// -> query
func User_login(c echo.Context) error {
	email := c.FormValue("email")
	pwd := c.FormValue("pwd")
	a_User := new(model.Users)
	a_User.Email = email
	err := model.DB.Debug().Find(&a_User).Error
	if err != nil {
		status = "nil"
		return response.SendResponse(c, 102, "incorrect email or password", status)
	}
	if a_User.Pwd != pwd {
		status = "nil"
		return response.SendResponse(c, 102, "incorrect email or password", status)
	}
	status = a_User.Name
	return response.SendResponse(c, 103, "login successfully", status)
}

func User_logout(c echo.Context) error {
	if status == "nil" {
		return response.SendResponse(c, 900, "didn' login", status)
	}
	status = "nil"
	return response.SendResponse(c, 900, "didn' login", status)
}

// {host}/user/security
func User_reset_pwd(c echo.Context) error {
	name := status
	if name == "nil" {
		return response.SendResponse(c, 900, "didn't login", status)
	}
	secA := c.FormValue("secA")
	newpwd := c.FormValue("newpwd")
	a_user := model.Users{}
	a_user.Name = name
	err := model.DB.Debug().Find(&a_user).Error
	if err != nil {
		logrus.Fatal("Unknown Error")
	}
	if a_user.SecA != secA {
		return response.SendResponse(c, 200, "incorrect answer", status)
	}
	a_user.Pwd = newpwd
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
		return response.SendResponse(c, 111, "nno User (FATAL)", status)
	}
	return response.SendResponse(c, 310, "getting information succeeded", a_user.LatestTime, status, a_user.Name)
}
