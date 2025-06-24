package service

import (
	"github.com/textures1245/payso-check-slip-backend/handler"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/repository"
)

type LogService interface {
	GetLogAll() ([]model.SureSureLog, error)
	GetLogByID(id int) ([]model.SureSureLog, error)
	CreateLog(log model.SureSureLog) (int, error)
	UpdateLog(log model.SureSureLog) error
	DeleteLog(id int) error
}

type logService struct {
	logHandler handler.LogHandler
}

func NewLogService(logHandler handler.LogHandler) LogService {
	return &logService{logHandler}
}

func (s *logService) GetLogAll() ([]model.SureSureLog, error) {
	return repository.GetLogAll()
}

func (s *logService) GetLogByID(id int) ([]model.SureSureLog, error) {
	return repository.GetLogByID(id)
}

func (s *logService) CreateLog(log model.SureSureLog) (int, error) {
	return repository.CreateLog(log)
}

func (s *logService) UpdateLog(log model.SureSureLog) error {
	return repository.UpdateLog(log)
}

func (s *logService) DeleteLog(id int) error {
	return repository.DeleteLog(id)
}
