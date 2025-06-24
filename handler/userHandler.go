package handler

type UserHandler interface {
}

type userHandler struct {
}

func NewUserHandler() userHandler {
	return userHandler{}
}
