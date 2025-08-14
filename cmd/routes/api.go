package routes

import (
	"github.com/JagTheFriend/Go-Chat-App/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	e  *echo.Echo
	db *database.DB
}

func NewAPI(e *echo.Echo) *API {
	return &API{
		e:  e,
		db: database.NewDB(),
	}
}

func (api *API) RegisterRoutes() {
	api.RegisterUserRoutes()
	api.RegisterMessageRoutes()
	api.RegisterWebsocket()
}

func (api *API) Start() {
	api.e.Use(middleware.Logger())
	api.e.Use(middleware.Recover())

	api.RegisterRoutes()

	api.e.Logger.Fatal(api.e.Start(":8080"))
}
