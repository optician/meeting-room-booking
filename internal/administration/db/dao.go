package db

import (
	"github.com/google/uuid"
	"github.com/optician/meeting-room-booking/internal/administration/models"
)

type DB interface {
	List() ([]models.RoomInfo, error)
	Update(models.RoomInfo) error
	Create(models.NewRoomInfo) (uuid.UUID, error)
	Delete(uuid.UUID) error
}

type impl struct{}

func New() DB {
	return impl{}
}

func (impl) List() ([]models.RoomInfo, error) {
	list := []models.RoomInfo{{Id: "123", Name: "Belyash", Capacity: 5, Office: "BC Utopia", Stage: 20, Labels: []string{"video", "projector"}}}
	return list, nil
}
func (impl) Update(models.RoomInfo) error {
	return nil
}
func (impl) Create(models.NewRoomInfo) (uuid.UUID, error) {
	id := uuid.New()
	return id, nil
}
func (impl) Delete(id uuid.UUID) error {
	return nil
}
