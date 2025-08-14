package routes

import (
	"github.com/JagTheFriend/Go-Chat-App/cmd/types"
	"github.com/labstack/echo/v4"
)

func (routes *API) RegisterUserRoutes() {
	route := routes.e.Group("/user")

	route.GET("/:id", func(c echo.Context) error {
		u := routes.db.Get(c.Param("id"))
		if u == nil {
			return c.JSON(404, nil)
		}
		return c.JSON(200, u)
	})

	route.POST("/new", func(c echo.Context) error {
		u := new(types.User)
		err := c.Bind(u)
		if err != nil {
			return err
		}
		routes.db.Set(u.ID, u)
		return c.JSON(200, u)
	})

	route.DELETE("/:id", func(c echo.Context) error {
		routes.db.Delete(c.Param("id"))
		return c.JSON(200, nil)
	})
}
