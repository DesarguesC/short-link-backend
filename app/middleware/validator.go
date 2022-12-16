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
	Name string `json:"name" validate:"required,excludesall=!@#$%^&*()_-{},ne=nil"`
	//ID     int    `json:"ID" validate:"required,min=0,max=10"`
	Email string `json:"email" validate:"required,contains=@[.*].com"`
	Pwd   string `json:"pwd" validate:"required,excludesall=!@#$%^&*()_-{}"`
	SecQ  string `json:"secq" validate:"required,excludesall=!@#$%^&*()_-{}"`
	SecA  string `jsong:"seca" validate:"required,excludesall=!@#$%^&*()_-{}"`
}

// 登录时的输入字符验证
type LoginStruct struct {
	Email string `json:"email" validate:"required,contains=@[.*].com"`
	Pwd   string `json:"pwd" validate:"required,excludesall=!@#$%^&*()_-{}"`
}

type SecurityStruct struct {
	Name    string `json:"name" validate:"required,excludesall=!@#$%^&*()_-{},ne=nil"`
	Email   string `json:"email" validate:"required,contains=@[.*].com"`
	Pwd_new string `json:"newpwd" validate:"required,excludesall=!@#$%^&*()_-{}"`
	SecA    string `json:"secA" validate:"required,excludesall=!@#$%^&*()_-{}"`
}

//// 在系统认为的已登录状态下的验证（验证state）
//type StateStruct struct {
//	State string `json:"state" validate:"required,excludesall=!@#$%^&*()_-{},ne=nil"`
//}

type ResetStruct struct {
	Name    string `json:"name" validate:"required,excludesall=!@#$%^&*()_-{},ne=nil"`
	Pwd_new string `json:"newpwd" validate:"required,excludesall=!@#$%^&*()_-{}"`
	Pwd_old string `json:"oldpwd" validate:"required,excludesall=!@#$%^&*()_-{}"`
}
