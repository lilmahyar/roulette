package room

import "context"

type Room struct{
	Id string
	Title string
	Description string
	//ProfileAvatar []byte
}

type Repository struct {
	GetRoomById(ctx context.Context, string id)
}


func GetRoomById(ctx context.Context, id string) ()