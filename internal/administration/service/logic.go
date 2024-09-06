package service

import "github.com/optician/meeting-room-booking/internal/administration/models"

type Logic interface {
	Create(room models.NewRoomInfo)

	Update(room models.RoomInfo)

	List() []models.RoomInfo

	Delete(id string)
}
