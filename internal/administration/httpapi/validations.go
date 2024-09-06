package httpapi

import (
	"encoding/json"
	"errors"
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

func validateRoomInfo(room *models.RoomInfo) (models.RoomInfo, error) {
	if room.Capacity < 1 {
		return *room, errors.New("room can't have 0 or less capacity")
	}
	if room.Id == "" {
		return *room, errors.New("room id can't be empty")
	}
	if room.Name == "" {
		return *room, errors.New("room name can't be empty")
	}
	if room.Office == "" {
		return *room, errors.New("room office can't be empty")
	}

	return *room, nil
}

func fromBytesRoom(stream io.Reader) (models.RoomInfo, error) {
	if room, err := deserializeRoom(stream); err != nil {
		return room, err
	} else {
		return validateRoomInfo(&room)
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

func validateNewRoomInfo(newRoom *models.NewRoomInfo) (models.NewRoomInfo, error) {
	if newRoom.Capacity < 1 {
		return *newRoom, errors.New("room can't have 0 or less capacity")
	}
	if newRoom.Name == "" {
		return *newRoom, errors.New("room name can't be empty")
	}
	if newRoom.Office == "" {
		return *newRoom, errors.New("room office can't be empty")
	}

	return *newRoom, nil
}

func fromBytesNewRoom(stream io.Reader) (models.NewRoomInfo, error) {
	if room, err := deserializeNewRoom(stream); err != nil {
		return room, err
	} else {
		return validateNewRoomInfo(&room)
	}
}
