package model
import "time"
// TODO: add new model
type Url struct{
	Id         int       `gorm:"type:uint;primaryKey autoincrement" form:"id" json:"id"`
	UserId     int       `gorm:"type:varchar(10) column:user_id" form:"user_id" json:"user_id"`
	Origin     string    `gorm:"type:varchar(200)" form:"origin" json:"origin"`
	Short      string    `gorm:"type:varchar(40)" form:"short" json:"short"`
	Comment    string    `gorm:"type:varchar(100)" form:"comment" json:"comment"`
	StartTime  time.Time `gorm:"type:datetime;autoCreateTime"`
	ExpireTime time.Time `gorm:"type:datetime"`
}