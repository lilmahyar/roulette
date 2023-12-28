package model

import "context"

type Repository interface {
	GetRoomById(ctx context.Context, id string) (Room, error)
	DeleteRoomById(ctx context.Context, id string) error
}

func GetRoomById(ctx context.Context, id string) (Room, error) {

}

func DeleteRoomById(ctx context.Context, id string) error {

}
