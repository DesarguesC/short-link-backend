package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/app/response"
	"go-svc-tpl/databases"
	"go-svc-tpl/model"
)

// url 的crud qwq

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
	if err != nil {
		logrus.Error("dbAdd err")
		return response.SendResponse(c, 400, "dbAdd err")
	}
	return response.SendResponse(c, 200, "create is ok")
}

func QueryUrl(c echo.Context) (err error) { //url details
	data := new(model.QueryInput)
	if err = c.Bind(data); err != nil {
		logrus.Error("Bind Fail")
	}
	resp, err := databases.QueryUrl((*data).Short)
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
	(*url).Comment = (*data).Comment
	(*url).ExpireTime = (*data).ExpireTime
	(*url).StartTime = (*data).StartTime
	(*url).Enable = true
	err = databases.UpdateUrl(url)
	if err != nil {
		return response.SendResponse(c, 400, "update failed")
	}
	return response.SendResponse(c, 200, "update succeed")
}

func DelUrl(c echo.Context) (err error) { //url details
	data := new(model.QueryInput)
	if err = c.Bind(data); err != nil {
		logrus.Error("Bind Fail")
	}
	err = databases.DelUrl((*data).Short)
	if err != nil {
		return response.SendResponse(c, 400, "Del failed")
	}
	return response.SendResponse(c, 200, "del succeed") //
}

func PauseUrl(c echo.Context) error { //
	data := new(model.DelInput)
	if err := c.Bind(data); err != nil {
		logrus.Error("Bind Failed")
	}
	err := databases.PauseUrl((*data).Short)
	if err != nil {
		return response.SendResponse(c, 400, "Pause failed")
	}
	return response.SendResponse(c, 200, "pause succeed") //
}
func ContinueUrl(c echo.Context) error {
	data := new(model.DelInput)
	if err := c.Bind(data); err != nil {
		logrus.Error("Bind Failed")
	}
	err := databases.ContinueUrl((*data).Short)
	if err != nil {
		return response.SendResponse(c, 400, "Continue failed")
	}
	return response.SendResponse(c, 200, "Continue succeed") //
}
