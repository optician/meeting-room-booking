package service

import (
	"github.com/optician/meeting-room-booking/internal/administration/models"
	"go.uber.org/zap"
)

type Logic interface {
	Create(room *models.NewRoomInfo) error

	Update(room *models.RoomInfo) error

	List() ([]models.RoomInfo, error) 

	Delete(id string) error
}

type impl struct {
	logger *zap.SugaredLogger
}

func Make(logger *zap.SugaredLogger) Logic {
	defer logger.Sync()

	return impl { 
		logger: logger,
	}
}

func (impl impl) Create(room *models.NewRoomInfo) error {
	impl.logger.Infof("recieved a new room %v", *room)
	return nil
}

func (impl impl) Update(room *models.RoomInfo) error {
	impl.logger.Infof("recieved an updated room %v", *room)
	return nil
}

func (impl impl) List() ([]models.RoomInfo, error) {
	list := []models.RoomInfo{{Id: "123", Name: "Belyash", Capacity: 5, Office: "BC Utopia", Stage: 20, Labels: []string{"video", "projector"}}}
	return list, nil
}

func (impl impl) Delete(id string) error {
	impl.logger.Infof("delete %v room", id)
	return nil
}
