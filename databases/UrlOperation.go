package databases

//gorm
import (
	"github.com/sirupsen/logrus"
	"go-svc-tpl/model"
)

func AddUrl(url *model.Url) error {
	tmp := *url
	err := model.DB.Debug().Create(&tmp).Error
	if err != nil {
		logrus.Error("sql addUrl fail")
	}
	return err
}

func QueryUrl(Id int) (*model.Url, error) {
	tmp := new(model.Url)
	tmp.Id = Id
	err := model.DB.Debug().Find(&tmp).Error

	if err != nil {
		logrus.Error("sql queryUrl fail")
	}
	return tmp, err // 返回整个结构体
}

func UpdateUrl(url *model.Url) error { //
	var tmp model.Url
	tmp = *url
	err := model.DB.Debug().Updates(&tmp).Error
	if err != nil {
		logrus.Error("sql update error")
	}
	return err
}

func DelUrl(Id int) error {
	tmp := new(model.Url)
	tmp.Id = Id
	err := model.DB.Debug().Delete(&tmp).Error
	if err != nil {
		logrus.Error("sql del fail")
	}
	return err
}

// Post
func PauseUrl(Id int) error {
	tmp, err := QueryUrl(Id)
	if err != nil {
		logrus.Error("pause not found")
	}
	(*tmp).Enable = false
	err = model.DB.Debug().Updates(tmp).Error
	if err != nil {
		logrus.Error("pause failed")
	}
	return err
}
func ContinueUrl(Id int) error {
	tmp, err := QueryUrl(Id)
	if err != nil {
		logrus.Error("continue Not found")
	}
	(*tmp).Enable = true
	err = model.DB.Debug().Updates(tmp).Error
	if err != nil {
		logrus.Error("continue failed")
	}
	return err
}
