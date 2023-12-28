package main

import (
	"github.com/labstack/echo"
	"roulette/cmd/database"
	"roulette/cmd/handlers/room"
)

func Main() {
	e := echo.New()

	database.ConnectToDatabase()

	room.RegisterHandler(e)

	e.Start(":5000")
}
