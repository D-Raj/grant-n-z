package api

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type RequestValidator struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &RequestValidator{validator: validator.New()}
}

func (gv *RequestValidator) Validate(i interface{}) error {
	return gv.validator.Struct(i)
}

func ValidateBody(c echo.Context, body interface{}) error {
	if err := c.Bind(body); err != nil {
		return err
	}
	if err := c.Validate(body); err != nil {
		return err
	}
	return nil
}