package databases

//gorm
import (
	"github.com/sirupsen/logrus"
	"go-svc-tpl/model"
)

func AddUrl(url *model.Url) {
	tmp := *url
	err := model.DB.Debug().Create(&tmp).Error
	if err != nil {
		logrus.Error("sql addUrl fail")
	}
}

func QueryUrl(Id int) *model.Url {
	tmp := new(model.Url)
	tmp.Id = Id
	err := model.DB.Debug().Find(&tmp).Error
	if err != nil {
		logrus.Error("sql queryUrl fail")
	}
	return tmp // 返回整个结构体
}

func UpdateUrl(url *model.Url) { //
	var tmp model.Url
	tmp = *url
	err := model.DB.Debug().Updates(&tmp).Error
	if err != nil {
		logrus.Error("sql update error")
	}
}

func DelUrl(Id int) {
	tmp := new(model.Url)
	tmp.Id = Id
	err := model.DB.Debug().Delete(&tmp).Error
	if err != nil {
		logrus.Error("sql del fail")
	}
}
