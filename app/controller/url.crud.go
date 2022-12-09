package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/app/response"
	"go-svc-tpl/databases"
	"go-svc-tpl/model"
)

// url 的crud qwq
func GenerateShortUrl(url *model.Url) { //生成短链接算法
	// 修改为短链接
}

func CreateUrl(c echo.Context) (err error) {
	data := new(model.CreateInput)
	if err = c.Bind(data); err != nil {
		logrus.Error("Bind Fail")
	}
	url := new(model.Url)
	(*url).Origin = (*data).Origin
	(*url).Comment = (*data).Comment
	(*url).Id = (*data).Id
	(*url).ExpireTime = (*data).ExpireTime
	(*url).StartTime = (*data).StartTime
	GenerateShortUrl(url)
	err = databases.AddUrl(url)
	return response.SendResponse(c, 102, "create is ok")
	// 改成response方法？
}

func QueryUrl(c echo.Context) (err error) { //url details
	var id int
	if err = c.Bind(id); err != nil {
		logrus.Error("Bind Fail")
	}
	resp, err := databases.QueryUrl(id)
	if err != nil {
		return response.SendResponse(c, 400, "query failed")
	}
	return response.SendResponse(c, 200, "query succeed", *resp)
}

// UpdateUrl  :post
func UpdateUrl(c echo.Context) (err error) { //url details
	data := new(model.UpdateInput)
	if err = c.Bind(data); err != nil {
		logrus.Error("Bind Fail")
	}
	url := new(model.Url)
	(*url).Origin = (*data).Origin
	(*url).Comment = (*data).Comment
	(*url).Id = (*data).Id
	(*url).ExpireTime = (*data).ExpireTime
	(*url).StartTime = (*data).StartTime
	err = databases.UpdateUrl(url)
	if err != nil {
		return response.SendResponse(c, 400, "update failed")
	}
	return response.SendResponse(c, 200, "update succeed")
}

func DelUrl(c echo.Context) (err error) { //url details
	var id int
	if err = c.Bind(id); err != nil {
		logrus.Error("Bind Failed")
	}
	err = databases.DelUrl(id)
	if err != nil {
		return response.SendResponse(c, 400, "Del failed")
	}
	return response.SendResponse(c, 200, "del succeed") //
}

// Post
func PauseUrl(c echo.Context) error { //
	var id int
	if err := c.Bind(id); err != nil {
		logrus.Error("Bind Failed")
	}
	err := databases.PauseUel(id)
	if err != nil {
		return response.SendResponse(c, 400, "Pause failed")
	}
	return response.SendResponse(c, 200, "pause succeed") //
}
