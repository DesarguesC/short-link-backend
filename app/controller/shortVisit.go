package controller

// 短链接访问
import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/app/response"
	"go-svc-tpl/model"
	"net/http"
	"time"
)

func Visit(c echo.Context) error {
	var short model.ShortUrl
	if err := c.Bind(short); err != nil {
		logrus.Error("bind error")
		return response.SendResponse(c, 400, "bind error")
	}
	url := new(model.Url)
	if err := model.DB.Debug().Where("short = ?", short).Find(url).Error; err != nil {
		logrus.Error("search failed")
		return response.SendResponse(c, 400, "search failed")
	}
	// 重定向
	if err := c.Redirect(http.StatusMovedPermanently, url.Origin); err != nil {
		logrus.Error("redirect failed")
		return response.SendResponse(c, 400, "redirect failed")
	}
	if time.Now().After(url.ExpireTime) { // 已过期
		logrus.Error("link expired")
	}
	return response.SendResponse(c, 200, "visit succeed")
}
