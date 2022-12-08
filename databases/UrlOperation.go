package databases

//gorm
import (
	"short-link-backend/model"

	"github.com/sirupsen/logrus"
)

func addUrl(url *model.Url) {
	tmp := *url
	err := model.DB.Debug().Create(&tmp).Error
	if err != nil {
		logrus.Error("sql addUrl fail")
	}
}

func queryUrl(Id int) *model.Url {
	tmp := new(model.Url)
	err := model.DB.Debug().Find(&tmp).Error
	if err != nil {
		logrus.Error("sql queryUrl fail")
	}
	return tmp
}

func updateUrl(url *model.Url) { //
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
