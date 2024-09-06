package httpapi

import (
	"strings"
	"testing"

	"github.com/optician/meeting-room-booking/internal/administration/models"
	"github.com/stretchr/testify/require"
)

func TestNewRoomDeserialization(t *testing.T) {
	json := `
	{
		"name":"Belyash",
		"capacity":5,
		"office":"BC Utopia",
		"stage":20,
		"labels":["video","projector"]
	}`
	expected := models.NewRoomInfo{
		Name:     "Belyash",
		Capacity: 5,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}

	actual, err := deserializeNewRoom(strings.NewReader(json))

	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestNewRoomCapacityValidationFailed(t *testing.T) {
	data := models.NewRoomInfo {
		Name:     "Belyash",
		Capacity: 0,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room can't have 0 or less capacity"
	_, err := models.ValidateNewRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestNewRoomNameValidationFailed(t *testing.T) {
	data := models.NewRoomInfo {
		Name:     "",
		Capacity: 10,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room name can't be empty"
	_, err := models.ValidateNewRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestNewRoomOfficeValidationFailed(t *testing.T) {
	data := models.NewRoomInfo {
		Name:     "Echpochmak",
		Capacity: 10,
		Office:   "",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room office can't be empty"
	_, err := models.ValidateNewRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestNewRoomValidationPassed(t *testing.T) {
	data := models.NewRoomInfo{
		Name:     "Belyash",
		Capacity: 1,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := data
	actual, err := models.ValidateNewRoomInfo(&data)

	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestRoomDeserialization(t *testing.T) {
	json := `
	  {
		"id": "123", 
		"name":"Belyash",
		"capacity":5,
		"office":
		"BC Utopia",
		"stage":20,
		"labels":["video","projector"]
	  }`
	expected := models.RoomInfo{
		Id:       "123",
		Name:     "Belyash",
		Capacity: 5,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}

	actual, err := deserializeRoom(strings.NewReader(json))

	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestRoomCapacityValidationFailed(t *testing.T) {
	data := models.RoomInfo{
		Id:       "123",
		Name:     "Belyash",
		Capacity: 0,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room can't have 0 or less capacity"
	_, err := models.ValidateRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestRoomIdValidationFailed(t *testing.T) {
	data := models.RoomInfo{
		Id:       "",
		Name:     "Belyash",
		Capacity: 5,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room id can't be empty"
	_, err := models.ValidateRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestRoomNameValidationFailed(t *testing.T) {
	data := models.RoomInfo{
		Id:       "123",
		Name:     "",
		Capacity: 5,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room name can't be empty"
	_, err := models.ValidateRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestRoomOfficeValidationFailed(t *testing.T) {
	data := models.RoomInfo{
		Id:       "123",
		Name:     "Matnakash",
		Capacity: 5,
		Office:   "",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room office can't be empty"
	_, err := models.ValidateRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestRoomValidationPassed(t *testing.T) {
	data := models.RoomInfo{
		Id:       "123",
		Name:     "Belyash",
		Capacity: 1,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := data
	actual, err := models.ValidateRoomInfo(&data)

	require.Nil(t, err)
	require.Equal(t, expected, actual)
}
