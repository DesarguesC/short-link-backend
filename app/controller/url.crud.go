package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/app/response"
	"go-svc-tpl/databases"
	"go-svc-tpl/model"
	"gorm.io/gorm"
)

// url 的crud qwq

func CreateUrl(c echo.Context) (err error) {
	data := new(model.CreateInput)
	if err = c.Bind(data); err != nil {
		logrus.Error(err)
		return response.SendResponse(c, 400, "Bind Failed")
	}

	url := new(model.Url)
	url.Origin = data.Origin
	url.Comment = data.Comment
	url.ExpireTime = data.ExpireTime
	url.StartTime = data.StartTime
	url.Enable = "able"
	err = model.DB.Debug().Where("origin = ?", (url).Origin).First(url).Error
	if err != gorm.ErrRecordNotFound {
		return response.SendResponse(c, 400, "have created")
	}
	err = model.DB.Debug().Create(url).Error
	//fmt.Println((*url).Id)
	if err != nil {
		logrus.Error(err)
		return response.SendResponse(c, 400, "dbAdd err")
	}
	GenerateShortUrl(url)
	err = model.DB.Debug().Updates(url).Error
	if err != nil {
		logrus.Error(err)
		return response.SendResponse(c, 400, "update error")
	}
	return response.SendResponse(c, 200, "create is ok", *url)
}

func QueryUrl(c echo.Context) (err error) { //url details
	data := new(model.QueryInput)
	if err = c.Bind(data); err != nil {
		logrus.Error(err)
		return response.SendResponse(c, 400, "Bind Fail")
	}
	resp, err := databases.QueryUrl((*data).Short)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.SendResponse(c, 400, "not found")
		}
		return response.SendResponse(c, 400, "sql queryUrl fail")
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
	url.Origin = data.Origin
	url.Comment = (data).Comment
	url.ExpireTime = data.ExpireTime
	url.StartTime = data.StartTime
	err = databases.UpdateUrl(url)
	if err != nil {
		return response.SendResponse(c, 400, "update failed")
	}
	return response.SendResponse(c, 200, "update succeed", *url) //
}

func DelUrl(c echo.Context) (err error) { //url details
	data := new(model.QueryInput)
	if err = c.Bind(data); err != nil {
		logrus.Error("Bind Fail")
	}
	err = databases.DelUrl(data.Short)
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
	err, resp := databases.PauseUrl(data.Short)
	if err != nil {
		return response.SendResponse(c, 400, "Pause failed", resp)
	}
	return response.SendResponse(c, 200, "pause succeed", resp) //
}
func ContinueUrl(c echo.Context) error {
	data := new(model.DelInput)
	if err := c.Bind(data); err != nil {
		logrus.Error(err)
		return response.SendResponse(c, 200, "Bind error") //
	}
	err, resp := databases.ContinueUrl(data.Short)
	if err != nil {
		return response.SendResponse(c, 400, "Continue failed", resp)
	}
	return response.SendResponse(c, 200, "Continue succeed", resp) //
}
