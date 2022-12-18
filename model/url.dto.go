package model

import "time"

type CreateInput struct { // 前端输入
	Origin     string    `gorm:"type:varchar(200)" form:"origin" json:"origin"`
	Comment    string    `gorm:"type:varchar(100)" form:"comment" json:"comment"`
	StartTime  time.Time `gorm:"type:datetime;autoCreateTime" json:"start-time"`
	ExpireTime time.Time `gorm:"type:datetime" json:"expire-time"`
}

type UpdateInput struct {
	Origin     string    `gorm:"type:varchar(200)"  json:"origin"`
	Comment    string    `gorm:"type:varchar(100)"  json:"comment"`
	StartTime  time.Time `gorm:"type:datetime;autoCreateTime" json:"start-time"`
	ExpireTime time.Time `gorm:"type:datetime" json:"expire-time"`
}

type DelInput struct {
	Short string `gorm:"type:varchar(40)" form:"short" json:"short"` //?
}
type QueryInput struct {
	Short string `gorm:"type:varchar(40)" form:"short" json:"short"` //?
}
type PauseInput struct {
	Short string `gorm:"type:varchar(40)" form:"short" json:"short"` //?
}
type ProfileInput struct {
	Id int `gorm:"type:uint;primaryKey autoincrement"  json:"id"`
	// period
}
type ProfileOutput struct {
	Time       time.Time `gorm:"type:datetime;autoCreateTime" json:"time"`
	AccessTime time.Time `gorm:"type:datetime;autoCreateTime" json:"access-time"`
}
