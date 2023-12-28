package main

import (
	"github.com/labstack/echo"
	"roulette/cmd/handlers/room"
)

func Main() {
	e := echo.New()

	room.RegisterHandler(e)

	e.Start(":5000")
}
