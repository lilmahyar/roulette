package model

type Room struct {
	ID          int `gorm:"autoIncrement"`
	Title       string
	Description string
	//ProfileAvatar []byte
}
