package middleware

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

type CustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

type RegisterStruct struct {
	Name string `json:"name"`
	//ID     int    `json:"ID" validate:"required,min=0,max=10"`
	Email string `json:"email" validate:"required,contains=@[.*].com"`
	Pwd   string `json:"pwd" validate:"required,excludesall=!@#$%^&*()_-{}"`
	SecQ  string `json:"secq" validate:"required,excludesall=!@#$%^&*()_-{}"`
	SecA  string `jsong:"seca" validate:"required,excludesall=!@#$%^&*()_-{}"`
}

type LoginStruct struct {
	Email string `json:"email" validate:"required,contains=@[.*].com"`
	Pwd   string `json:"pwd" validate:"required,excludesall=!@#$%^&*()_-{}"`
}
