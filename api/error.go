package api

import (
	"github.com/labstack/echo"
	"net/http"
)

type Error struct {
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

func ErrorHandler(err error, c echo.Context) {
	var response Error
	var code int
	var message string

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		switch code {
		case http.StatusBadRequest:
			message = "bad request."
		case http.StatusUnauthorized:
			message = "unauthorized."
		case http.StatusForbidden:
			message = "forbidden."
		case http.StatusNotFound:
			message = "not found."
		case http.StatusMethodNotAllowed:
			message = "method not allowed."
		case http.StatusConflict:
			message = "conflict resource."
		case http.StatusUnprocessableEntity:
			message = "unprocessable entity."
		case http.StatusInternalServerError:
			message = "internal server error."
		}

		response = Error {
			Message: message,
			Detail:  he.Message,
		}
	}

	c.JSON(code, response)
	c.Logger().Error(err)
}
