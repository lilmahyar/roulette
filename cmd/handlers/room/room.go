package room

import (
	"github.com/labstack/echo"
	"roulette/cmd/repository/room"
)

func RegisterHandler(e *echo.Echo) {
	e.GET("/rooms/:id", getRoom)
	e.POST("/rooms/", createNewRoom)
	e.PUT("/rooms/:id", editRoomInfo)
	e.DELETE("/rooms/:id", deleteRoom)
}

func createNewRoom(ctx echo.Context) error {

}

func editRoomInfo(ctx echo.Context) error {

}

func getRoom(ctx echo.Context) error {

	params :=
		room.Repository.GetRoomById(ctx, )
}

func deleteRoom(e echo.Context) error {

}
