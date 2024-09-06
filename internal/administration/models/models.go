package models

type RoomInfo struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Capacity int      `json:"capacity"`
	Office   string   `json:"office"`
	Stage    int      `json:"stage"`
	Labels   []string `json:"labels"`
}

type NewRoomInfo struct {
	Name     string   `json:"name"`
	Capacity int      `json:"capacity"`
	Office   string   `json:"office"`
	Stage    int      `json:"stage"`
	Labels   []string `json:"labels"`
}
