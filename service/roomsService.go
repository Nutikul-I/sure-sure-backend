package service

import (
	"github.com/textures1245/payso-check-slip-backend/handler"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/repository"
)

type RoomsService interface {
	GetAllRooms() ([]model.SureSureRoom, error)
	GetRoomByID(id int) ([]model.SureSureRoom, error)
	CreateRoom(room model.SureSureRoom) (int, error)
	UpdateRoom(room model.SureSureRoom) error
	DeleteRoom(id int) error
	HowTo(id int, user_id string) error
}

type roomsService struct {
	roomHandler handler.RoomHandler
}

func NewRoomsService(roomsHandler handler.RoomsHandler) RoomsService {
	return &roomsService{roomsHandler}
}

func (s *roomsService) GetAllRooms() ([]model.SureSureRoom, error) {
	return repository.GetAllRooms()
}

func (s *roomsService) GetRoomByID(id int) ([]model.SureSureRoom, error) {
	return repository.GetRoomByID(id)
}

func (s *roomsService) CreateRoom(room model.SureSureRoom) (int, error) {
	return repository.CreateRoom(room)
}

func (s *roomsService) UpdateRoom(room model.SureSureRoom) error {
	return repository.UpdateRoom(room)
}

func (s *roomsService) DeleteRoom(id int) error {
	return repository.DeleteRoom(id)
}

func (s *roomsService) HowTo(id int, user_id string) error {
	return repository.HowTo(id, user_id)
}
