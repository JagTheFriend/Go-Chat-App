package routes

import (
	"encoding/json"
	"fmt"

	"github.com/JagTheFriend/Go-Chat-App/database"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

func (routes *API) RegisterWebsocket() {
	routes.e.GET("/ws", routes.websocketRoute)
}
func (routes *API) websocketRoute(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer func(ws *websocket.Conn) {
		if err := ws.Close(); err != nil {
			c.Logger().Error(err)
		}
	}(ws)

	for {
		err = ws.WriteMessage(websocket.TextMessage, []byte("Hello world"))
		if err != nil {
			c.Logger().Error(err)
		}
		for data := range database.DbChan {
			fmt.Println(data)
			d := data.Data
			d["type"] = data.Type
			jsonData, err := json.Marshal(d)
			if err != nil {
				c.Logger().Error(err)
				return err
			}
			err = ws.WriteMessage(websocket.TextMessage, jsonData)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
