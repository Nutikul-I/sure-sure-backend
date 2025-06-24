package handler

type RoomHandler interface {
}

type roomHandler struct {
}

func NewRoomHandler() roomHandler {
	return roomHandler{}
}
