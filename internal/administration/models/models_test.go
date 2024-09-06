package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRoomCapacityValidationFailed(t *testing.T) {
	data := NewRoomInfo {
		Name:     "Belyash",
		Capacity: 0,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room can't have 0 or less capacity"
	_, err := ValidateNewRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestNewRoomNameValidationFailed(t *testing.T) {
	data := NewRoomInfo {
		Name:     "",
		Capacity: 10,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room name can't be empty"
	_, err := ValidateNewRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestNewRoomOfficeValidationFailed(t *testing.T) {
	data := NewRoomInfo {
		Name:     "Echpochmak",
		Capacity: 10,
		Office:   "",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room office can't be empty"
	_, err := ValidateNewRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestNewRoomValidationPassed(t *testing.T) {
	data := NewRoomInfo{
		Name:     "Belyash",
		Capacity: 1,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := data
	actual, err := ValidateNewRoomInfo(&data)

	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestRoomCapacityValidationFailed(t *testing.T) {
	data := RoomInfo{
		Id:       "123",
		Name:     "Belyash",
		Capacity: 0,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room can't have 0 or less capacity"
	_, err := ValidateRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestRoomIdValidationFailed(t *testing.T) {
	data := RoomInfo{
		Id:       "",
		Name:     "Belyash",
		Capacity: 5,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room id can't be empty"
	_, err := ValidateRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestRoomNameValidationFailed(t *testing.T) {
	data := RoomInfo{
		Id:       "123",
		Name:     "",
		Capacity: 5,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room name can't be empty"
	_, err := ValidateRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestRoomOfficeValidationFailed(t *testing.T) {
	data := RoomInfo{
		Id:       "123",
		Name:     "Matnakash",
		Capacity: 5,
		Office:   "",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := "room office can't be empty"
	_, err := ValidateRoomInfo(&data)

	require.EqualError(t, err, expected)
}

func TestRoomValidationPassed(t *testing.T) {
	data := RoomInfo{
		Id:       "123",
		Name:     "Belyash",
		Capacity: 1,
		Office:   "BC Utopia",
		Stage:    20,
		Labels:   []string{"video", "projector"},
	}
	expected := data
	actual, err := ValidateRoomInfo(&data)

	require.Nil(t, err)
	require.Equal(t, expected, actual)
}
