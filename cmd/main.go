package main

import (
	"github.com/JagTheFriend/Go-Chat-App/cmd/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	api := routes.NewAPI(e)
	api.Start()
}
