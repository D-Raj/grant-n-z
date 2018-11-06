package controller

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/handler"
	"github.com/tomoyane/grant-n-z/infra"
	"net/http"
)

func PostService(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	errAuth := di.ProvideTokenService.VerifyToken(c, token, "")
	if errAuth != nil {
		return echo.NewHTTPError(errAuth.Code, errAuth)
	}

	service := new(entity.Service)

	if err = c.Bind(service); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err = c.Validate(service); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	serviceData := di.ProvideServiceService.InsertService(service)

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/services/" + serviceData.Uuid.String())
	return c.JSON(http.StatusCreated, serviceData)
}

func GetService(c echo.Context) (err error) {
	serviceData := di.ProvideServiceService.GetAll()

	return c.JSON(http.StatusOK, serviceData)
}