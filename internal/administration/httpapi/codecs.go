package httpapi

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/optician/meeting-room-booking/internal/administration/models"
)

func deserializeRoom(stream io.Reader) (models.RoomInfo, error) {
	room := &models.RoomInfo{}
	if err := json.NewDecoder(stream).Decode(room); err != nil {
		return *room, fmt.Errorf("can't deserialize RoomInfo: %w", err)
	} else {
		return *room, nil
	}
}

func fromBytesRoom(stream io.Reader) (models.RoomInfo, error) {
	if room, err := deserializeRoom(stream); err != nil {
		return room, err
	} else {
		return models.ValidateRoomInfo(&room)
	}
}

func deserializeNewRoom(stream io.Reader) (models.NewRoomInfo, error) {
	room := &models.NewRoomInfo{}
	if err := json.NewDecoder(stream).Decode(room); err != nil {
		return *room, fmt.Errorf("can't deserialize NewRoomInfo: %w", err)
	} else {
		return *room, nil
	}
}

func fromBytesNewRoom(stream io.Reader) (models.NewRoomInfo, error) {
	if room, err := deserializeNewRoom(stream); err != nil {
		return room, err
	} else {
		return models.ValidateNewRoomInfo(&room)
	}
}
