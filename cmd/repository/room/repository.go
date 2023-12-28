package room

import (
	"context"
	"roulette/cmd/database"
	"roulette/cmd/model"
)

type Repository interface {
	GetRoomById(ctx context.Context, id string) model.Room
	DeleteRoomById(ctx context.Context, id string)
}

func GetRoomById(ctx context.Context, id string) model.Room {
	var room model.Room

	database.DbInstance.First(&room, id)

	return room
}

func DeleteRoomById(ctx context.Context, id string) {
	database.DbInstance.Delete(&model.Room{}, "id")
}
