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
	short := c.Param("hash")
	url := new(model.Url)
	if err := model.DB.Debug().Where("short = ?", short).Find(url).Error; err != nil {
		logrus.Error("search failed")
		return response.SendResponse(c, 400, "search failed")
	}
	//已冻结
	if url.Enable == false {
		logrus.Error("The link paused")
		return response.SendResponse(c, 400, "The link paused")
	}
	// 已过期
	if time.Now().After(url.ExpireTime) {
		logrus.Error("the link expired")
		return response.SendResponse(c, 400, "The link expired")
	}
	// 重定向
	if err := c.Redirect(http.StatusMovedPermanently, url.Origin); err != nil {
		logrus.Error("redirect failed")
		return response.SendResponse(c, 400, "redirect failed")
	}

	return response.SendResponse(c, 200, "visit succeed")
}
