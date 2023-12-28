package room

import "github.com/labstack/echo"

func RegisterHandler(e *echo.Echo) {
	e.GET("/rooms/:id", getRoom)
	e.POST("/rooms/", createNewRoom)
	e.PUT("/rooms/:id", editRoomInfo)
	e.DELETE("/rooms/:id", deleteRoom)
}

func createNewRoom(e echo.Context) error {

}

func editRoomInfo(e echo.Context) error {

}

func getRoom(e echo.Context) error {

}

func deleteRoom(e echo.Context) error {

}
