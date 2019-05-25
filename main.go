package main

import (
	"net/http"

	"github.com/ajangi/nardoon/controllers"
	"github.com/ajangi/nardoon/utils"
	"github.com/labstack/echo"
)

func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		emptyData := utils.ResponseData{}
		notFoundMessage := utils.ResponseMessages{utils.GetMessageByKey(utils.NotFoundErrorMessageKey)}
		errorResponse := utils.ResponseApi{Result: "ERROR", Data: emptyData, Messages: notFoundMessage, StatusCode: http.StatusNotFound}
		return c.JSON(http.StatusNotFound, errorResponse)
	}
	e := echo.New()
	e.GET("/healthy", func(c echo.Context) error {
		emptyData := utils.ResponseData{}
		healthyMessage := utils.ResponseMessages{utils.GetMessageByKey(utils.HealthyMessageKey)}
		healthyResponse := utils.ResponseApi{Result: "SUCCESS", Data: emptyData, Messages: healthyMessage, StatusCode: http.StatusOK}
		return c.JSON(http.StatusOK, healthyResponse)
	})
	usersGroup := e.Group("/users")
	usersGroup.POST("", controllers.RegisterUser)
	e.Logger.Fatal(e.Start(":1323"))
}
