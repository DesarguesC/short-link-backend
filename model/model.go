package model

import "time"

// TODO:  add new model
type Url struct {
	Id int `gorm:"type:uint;primaryKey autoincrement" form:"id" json:"id"`
	//UserId     int       `gorm:"type:varchar(10) column:user_id" form:"user_id" json:"user_id"`
	Origin     string    `gorm:"type:varchar(200)" form:"origin" json:"origin"`
	Short      string    `gorm:"type:varchar(40)" form:"short" json:"short"` //?
	Comment    string    `gorm:"type:varchar(100)" form:"comment" json:"comment"`
	StartTime  time.Time `gorm:"type:datetime;autoCreateTime" json:"start-time"`
	ExpireTime time.Time `gorm:"type:datetime" json:"expire-time"`
	Enable     bool      `gorm:"type:bool" json:"enable"`
}

type Users struct {
	Name  string `gorm:"type:varchar(20)"`
	Email string `gorm:"type:varchar(50)"`
	Pwd   string `gorm:"type:varchar(90)"`
	SecQ  string `gorm:"type:varchar(100)"`
	SecA  string `gorm:"type:varchar(100)"`
}

// dto 与前端交互
type CreateInput struct { // 前端输入
	Id         int       `gorm:"type:uint;primaryKey AUTO_INCREMENT" form:"id" json:"id"`
	Origin     string    `gorm:"type:varchar(200)" form:"origin" json:"origin"`
	Comment    string    `gorm:"type:varchar(100)" form:"comment" json:"comment"`
	StartTime  time.Time `gorm:"type:datetime;autoCreateTime" json:"start-time"`
	ExpireTime time.Time `gorm:"type:datetime" json:"expire-time"`
}

type UpdateInput struct {
	Id         int       `gorm:"type:uint;primaryKey AUTO_INCREMENT"  json:"id"`
	Origin     string    `gorm:"type:varchar(200)"  json:"origin"`
	Comment    string    `gorm:"type:varchar(100)"  json:"comment"`
	StartTime  time.Time `gorm:"type:datetime;autoCreateTime" json:"start-time"`
	ExpireTime time.Time `gorm:"type:datetime" json:"expire-time"`
}

type ProfileInput struct {
	Id int `gorm:"type:uint;primaryKey autoincrement"  json:"id"`
	// period
}
type ProfileOutput struct {
	Time       time.Time `gorm:"type:datetime;autoCreateTime" json:"time"`
	AccessTime time.Time `gorm:"type:datetime;autoCreateTime" json:"access-time"`
}
