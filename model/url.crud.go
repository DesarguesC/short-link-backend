package model

import (
	"net/http"
	"short-link-backend/databases"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// url 的crud qwq
func GenerateShortUrl(url *Url) { //生成短链接算法
	// 修改为短链接
}

func CreateUrl(c echo.Context) (err error) {
	url := new(Url)
	if err = c.Bind(url); err != nil {
		logrus.Error("Bind Fail")
	}
	GenerateShortUrl(url)
	databases.addUrl(url)
	return c.JSON(http.StatusOK, url.Id)
	// 改成response方法？
}

func QueryUrl(c echo.Context) (err error) { //url details
	var id int
	if err = c.Bind(id); err != nil {
		logrus.Error("Bind Fail")
	}
	resp := databases.queryUrl(id)
	return c.JSON(http.StatusOK, resp)
}

func UpdateUrl(c echo.Context) (err error) { //url details
	url := new(Url)
	if err = c.Bind(url); err != nil {
		logrus.Error("Bind Fail")
	}
	databases.updateUrl(url)
	return c.JSON(http.StatusOK, nil) //?
}

func DelUrl(c echo.Context) (err error) { //url details
	var id int
	if err = c.Bind(id); err != nil {
		logrus.Error("Bind Fail")
	}
	databases.DelUrl(id)
	return c.JSON(http.StatusOK, nil) //?
}
