package main

import (
	"net/http"

	"github.com/ajangi/nardoon/utils"
	"github.com/labstack/echo"
)

func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		emptyData := utils.ResponseData{}
		notFoundMessage := utils.ResponseMessages{utils.GetErrorByKey(utils.NotFoundErrorMessageKey)}
		errorResponse := utils.ResponseApi{Result: "ERROR", Data: emptyData, Messages: notFoundMessage, StatusCode: http.StatusNotFound}
		return c.JSON(http.StatusNotFound, errorResponse)
	}
	e := echo.New()
	e.GET("/healthy", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
