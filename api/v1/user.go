package v1

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/api"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/infra"
	"net/http"
	"strings"
)

func PostUser(c echo.Context) (err error) {
	body := new(entity.UserReq)
	if err := api.ValidateBody(c, body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request parameter is invalid.")
	}

	user, err, code := api.NewUserService().
	if err != nil {
		return echo.NewHTTPError(code, err.Error())
	}

	c.Response().Header().Add("Location", infra.GetHostName() + "/v1/users/" + user.Uuid.String())
	return c.JSON(http.StatusCreated, map[string]string {"message": "user creation succeeded."})
}

func PutUser(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	column := c.QueryParam("column")
	errAuth := di.ProvideTokenService.VerifyToken(c, token, "")

	if errAuth != nil {
		return echo.NewHTTPError(errAuth.Code, errAuth)
	}

	user := new(entity.User)
	if !strings.Contains(column, "username") &&
		!strings.EqualFold(column, "email") && !strings.EqualFold(column, "password") {

		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, handler.BadRequest(""))
	}

	errUser := di.ProvideUserService.PutUserColumnData(user, column)
	if errUser != nil {
		return echo.NewHTTPError(errUser.Code, errUser)
	}

	return c.JSON(http.StatusOK, map[string]string {"message": "ok."})
}
