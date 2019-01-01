package v1

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/tomoyane/grant-n-z/domain/entity"
	"net/http"
)

func PostToken(c echo.Context) (err error) {
	user := new(entity.User)
	fmt.Println(user.Email)
	if err = c.Bind(user); err != nil {
		return err
	}

	return c.JSON(http.StatusOK,  map[string]string {"token": "aaa"})
}