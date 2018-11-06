package v1

import (
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"net/http"
)

func PostRole(c echo.Context) (err error) {
	role := new(entity.Role)
	if err := c.Bind(role); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, "a")
}