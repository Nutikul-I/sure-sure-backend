package handler

type RoomsHandler interface {
}

type roomsHandler struct {
}

func NewRoomsHandler() roomsHandler {
	return roomsHandler{}
}
