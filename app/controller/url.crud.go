package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/databases"
	"go-svc-tpl/model"
	"net/http"
	"strconv"
)

// url 的crud qwq
func GenerateShortUrl(url *model.Url) { //生成短链接算法
	// 修改为短链接
}

func CreateUrl(c echo.Context) (err error) {
	url := new(model.Url)
	if err = c.Bind(url); err != nil {
		logrus.Error("Bind Fail")
	}
	GenerateShortUrl(url)
	databases.AddUrl(url)
	return c.JSON(http.StatusOK, url.Id)
	// 改成response方法？
}

func QueryUrl(c echo.Context) (err error) { //url details
	var id int
	if err = c.Bind(id); err != nil {
		logrus.Error("Bind Fail")
	}
	resp := databases.QueryUrl(id)
	return c.JSON(http.StatusOK, resp)
}

// UpdateUrl  :post
func UpdateUrl(c echo.Context) (err error) { //url details
	url := new(model.Url)
	url.Id, _ = strconv.Atoi(c.FormValue("Id"))
	url.Origin = c.FormValue("origin")
	databases.UpdateUrl(url)
	return c.JSON(http.StatusOK, nil) //
}

func DelUrl(c echo.Context) (err error) { //url details
	var id int
	if err = c.Bind(id); err != nil {
		logrus.Error("Bind Failed")
	}
	databases.DelUrl(id)
	return c.JSON(http.StatusOK, nil) //
}

// Post
func PauseUrl(c echo.Context) error { //
	var id int
	if err := c.Bind(id); err != nil {
		logrus.Error("Bind Failed")
	}
	databases.PauseUel(id)
	return c.JSON(http.StatusOK, nil) //
}
