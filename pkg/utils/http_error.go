package utils

import (
	"net/http"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/labstack/echo"
)

func HttpError(c echo.Context, err error) error {
	var statusCode int
	if e, ok := err.(*response_mapper.ResponseError); ok {
		statusCode = response_mapper.StatusErrorMapping(e.Code)
	}
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}

	return c.JSON(statusCode, response_mapper.RenderStruct(statusCode, err))
}
