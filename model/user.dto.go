package model

type RegisterInput struct {
	Id    int    `gorm:"type:uint;primaryKey autoincrement=1000" form:"id" json:"id"`
	Name  string `gorm:"type:varchar(100)" form:"name" json:"name"`
	Email string `gorm:"type:varchar(100)" form:"email" json:"email"`
	Pwd   string `gorm:"type:varchar(100)" form:"pwd" json:"pwd"`
	SecQ  string `gorm:"type:varchar(100)" form:"secq" json:"secq"`
	SecA  string `gorm:"type:varchar(100)" form:"seca" json:"seca"`
}

type LoginInput struct {
	Email string `gorm:"type:varchar(100)" form:"email" json:"email"`
	//Name  string `gorm:"type:varchar(20)" form:"name" json:"name"`
	Pwd string `gorm:"type:varchar(100)" form:"pwd" json:"pwd"`
}

type SecureResetPwdInput struct {
	Name    string `gorm:"type:varchar(100)" form:"name" json:"name"`
	Email   string `gorm:"type:varchar(100)" form:"email" json:"email"`
	SecA    string `gorm:"type:varchar(100)" form:"seca" json:"seca"`
	Pwd_new string `gorm:"type:varchar(100)" form:"newpwd" json:"newpwd"`
}

type ResetPwdInput struct {
	Email  string `gorm:"type:varchar(100)" form:"email" json:"email"`
	Name   string `gorm:"type:varchar(100)" form:"name" json:"name"`
	Pwdold string `gorm:"type:varchar(100)" form:"oldpwd" json:"oldpwd"`
	Pwdnew string `gorm:"type:varchar(100)" form:"newpwd" json:"newpwd"`
}
