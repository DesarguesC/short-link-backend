package middleware

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

type CustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

// 注册时的输入字符验证
type RegisterStruct struct {
	Name string `form:"name" json:"name" query:"name" validate:"excludesall=!@#$%^&*()_-{},ne=nil"`
	//ID     int    `json:"ID" validate:"required,min=0,max=10"`
	Email string `form:"email" json:"email" query:"email" validate:"contains=@[.]*.com"`
	Pwd   string `form:"pwd" json:"pwd" query:"pwd" validate:"excludesall=!@#$%^&*()_-{}"`
	SecQ  string `form:"secQ" json:"secQ" query:"secQ" validate:"excludesall=!@#$%^&*()_-{}"`
	SecA  string `form:"secA" json:"secA" query:"secA" validate:"excludesall=!@#$%^&*()_-{}"`
}

// 登录时的输入字符验证
type LoginStruct struct {
	Email string `form:"email" json:"email" query:"email" validate:"contains=@[.]*.com"`
	//Name  string `json:"name" validate:"required,excludesall=!@#$%^&*()_-{},ne=nil"`
	Pwd string `form:"pwd" json:"pwd" query:"pwd" validate:"excludesall=!@#$%^&*()_-{}"`
}

type SecurityStruct struct {
	Name    string `form:"name" json:"name" query:"name" validate:"excludesall=!@#$%^&*()_-{},ne=nil"`
	Email   string `form:"email" json:"email" query:"email" validate:"contains=@[.*].com"`
	Pwd_new string `form:"newpwd" json:"newpwd" query:"newpwd" validate:"excludesall=!@#$%^&*()_-{}"`
	SecA    string `form:"secA" json:"secA" query:"secA" validate:"excludesall=!@#$%^&*()_-{}"`
}

//// 在系统认为的已登录状态下的验证（验证state）
//type StateStruct struct {
//	State string `json:"state" validate:"required,excludesall=!@#$%^&*()_-{},ne=nil"`
//}

type ResetStruct struct {
	Name    string `form:"name" json:"name" validate:"required,excludesall=!@#$%^&*()_-{},ne=nil"`
	Pwd_new string `form:"newpwd" json:"newpwd" validate:"required,excludesall=!@#$%^&*()_-{}"`
	Pwd_old string `form:"oldpwd" json:"oldpwd" validate:"required,excludesall=!@#$%^&*()_-{}"`
}
