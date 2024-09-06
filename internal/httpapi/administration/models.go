package administration

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type RoomInfo struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Capacity int      `json:"capacity"`
	Office   string   `json:"office"`
	Stage    int      `json:"stage"`
	Labels   []string `json:"labels"`
}

func deserializeRoom(stream io.Reader) (RoomInfo, error) {
	room := &RoomInfo{}
	if err := json.NewDecoder(stream).Decode(room); err != nil {
		return *room, fmt.Errorf("can't deserialize RoomInfo: %w", err)
	} else {
		return *room, nil
	}
}

func validateRoomInfo(room *RoomInfo) (RoomInfo, error) {
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

func fromBytesRoom(stream io.Reader) (RoomInfo, error) {
	if room, err := deserializeRoom(stream); err != nil {
		return room, err
	} else {
		return validateRoomInfo(&room)
	}
}

type NewRoomInfo struct {
	Name     string   `json:"name"`
	Capacity int      `json:"capacity"`
	Office   string   `json:"office"`
	Stage    int      `json:"stage"`
	Labels   []string `json:"labels"`
}

func deserializeNewRoom(stream io.Reader) (NewRoomInfo, error) {
	room := &NewRoomInfo{}
	if err := json.NewDecoder(stream).Decode(room); err != nil {
		return *room, fmt.Errorf("can't deserialize NewRoomInfo: %w", err)
	} else {
		return *room, nil
	}
}

func validateNewRoomInfo(newRoom *NewRoomInfo) (NewRoomInfo, error) {
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

func fromBytesNewRoom(stream io.Reader) (NewRoomInfo, error) {
	if room, err := deserializeNewRoom(stream); err != nil {
		return room, err
	} else {
		return validateNewRoomInfo(&room)
	}
}
