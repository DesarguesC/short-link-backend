package middleware

import (
	"github.com/go-playground/validator/v10"
)

func (c *CustomValidator) Validate(i interface{}) error {
	c.ToInit()
	return c.validate.Struct(i)
}

func (c *CustomValidator) ToInit() {
	c.once.Do(func() {
		c.validate = validator.New()
	})
}
