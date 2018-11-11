package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/api"
	"github.com/tomoyane/grant-n-z/domain/entity"
)

func PostGroup(c echo.Context) (err error) {
	body := new(entity.Group)
	if err := api.ValidateBody(c, body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request parameter is invalid.")
	}

	group, err, code := api.NewGroupService().Insert(body)
	if err != nil {
		return echo.NewHTTPError(code, err.Error())
	}

	return c.JSON(code, group)
}

func GetGroups(c echo.Context) (err error) {
	groups, err, code := api.NewGroupService().GetAll()
	if err != nil {
		return echo.NewHTTPError(code, err.Error())
	}

	return c.JSON(code, groups)
}

func GetGroup(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, "path parameter is only number.")
	}

	group, err, code := api.NewGroupService().GetById(id)
	if err != nil {
		return echo.NewHTTPError(code, err.Error())
	}

	return c.JSON(code, group)
}
