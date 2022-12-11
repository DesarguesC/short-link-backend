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

func QueryUrl(short string) (*model.Url, error) {
	tmp := new(model.Url)
	tmp.Short = short
	//where ()可删？
	err := model.DB.Debug().Where("short = ?", short).Find(&tmp).Error
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

func DelUrl(short string) error {
	tmp := new(model.Url)
	tmp.Short = short
	err := model.DB.Debug().Where("short = ?", short).Delete(&tmp).Error
	if err != nil {
		logrus.Error("sql del fail")
	}
	return err
}

// Post
func PauseUrl(short string) error {
	tmp, err := QueryUrl(short)
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
func ContinueUrl(short string) error {
	tmp, err := QueryUrl(short)
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
