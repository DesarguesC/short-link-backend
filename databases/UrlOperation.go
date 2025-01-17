package databases

//gorm
import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-svc-tpl/model"
)

func QueryUrl(short string) (*model.Url, error) {
	tmp := new(model.Url)
	tmp.Short = short
	fmt.Println(short)
	err := model.DB.Debug().Where("short = ?", short).First(&tmp).Error
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println()
	return tmp, err // 返回整个结构体
}

func UpdateUrl(url *model.Url) error { //
	var tmp model.Url
	tmp = *url
	err := model.DB.Debug().Where("origin=?", tmp.Origin).Updates(&tmp).Error
	if err != nil {
		logrus.Error(err)
	}
	return err
}

func DelUrl(short string) error {
	tmp := new(model.Url)
	tmp.Short = short
	err := model.DB.Debug().Where("short = ?", short).Delete(tmp).Error
	if err != nil {
		logrus.Error(err)
	}
	return err
}

// Post
// updates 无法更新0值（flase） qwq 查半天
func PauseUrl(short string) (error, *model.Url) {
	tmp, err := QueryUrl(short)
	if err != nil {
		logrus.Error(err)
	}
	tmp.Enable = "unable"
	err = model.DB.Debug().Where("short = ?", tmp.Short).Updates(&tmp).Error
	if err != nil {
		logrus.Error(err)
	}
	return err, tmp
}
func ContinueUrl(short string) (error, *model.Url) {
	tmp, err := QueryUrl(short)
	if err != nil {
		logrus.Error(err)
	}
	tmp.Enable = "able"
	err = model.DB.Debug().Where("short = ?", short).Updates(&tmp).Error
	if err != nil {
		logrus.Error(err)
	}
	return err, tmp
}
