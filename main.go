package main

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/api"
	"github.com/tomoyane/grant-n-z/api/v1"
	"github.com/tomoyane/grant-n-z/infra"
)

func init() {
	infra.Init()
}

func main() {
	e := echo.New()
	e.HTTPErrorHandler = api.ErrorHandler
	e.Validator = api.NewValidator()

	e.POST("/v1/groups", v1.PostGroup)
	e.POST("/v1/groups", v1.GetGroups)
	e.POST("/v1/groups/:id", v1.GetGroup)

	e.Logger.Fatal(e.Start(":8080"))
}