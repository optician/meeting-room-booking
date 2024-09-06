package db

import "github.com/optician/meeting-room-booking/internal/administration/models"

type DB interface {
	List() ([]models.RoomInfo, error)
	Update(models.RoomInfo) error
	Create(models.NewRoomInfo) (string, error)
	Delete(string) error
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
func (impl) Create(models.NewRoomInfo) (string, error) {
	return "bla-bla", nil
}
func (impl) Delete(string) error {
	return nil
}
