package models

import "errors"

type RoomInfo struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Capacity int      `json:"capacity"`
	Office   string   `json:"office"`
	Stage    int      `json:"stage"`
	Labels   []string `json:"labels"`
}

func ValidateRoomInfo(room *RoomInfo) (RoomInfo, error) {
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

type NewRoomInfo struct {
	Name     string   `json:"name"`
	Capacity int      `json:"capacity"`
	Office   string   `json:"office"`
	Stage    int      `json:"stage"`
	Labels   []string `json:"labels"`
}

func ValidateNewRoomInfo(newRoom *NewRoomInfo) (NewRoomInfo, error) {
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
