package httpapi

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/optician/meeting-room-booking/internal/administration/models"
	"github.com/optician/meeting-room-booking/internal/administration/service"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var logger = zap.NewExample().Sugar()
var logic service.Logic = logicStub{}

func executeRequest(req *http.Request, s *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestListRoomsSuccessfully(t *testing.T) {
	// Create a New Server Struct
	r := chi.NewRouter()
	r.Route("/", Make(&logic, logger))

	req, _ := http.NewRequest("GET", "/rooms", nil)

	// Execute Request
	response := executeRequest(req, r)

	checkResponseCode(t, http.StatusOK, response.Code)
	expected := fmt.Sprintf(
		`[{"id":"%v","name":"Belyash","capacity":5,"office":"BC Utopia","stage":20,"labels":["video","projector"]}]`,
		stubId,
	)
	require.Equal(t, expected, response.Body.String())
}

func TestCreateRoomSuccessfully(t *testing.T) {
	// Create a New Server Struct
	r := chi.NewRouter()
	r.Route("/", Make(&logic, logger))

	// Create a New Request
	json := `
	{
		"name":"Belyash",
		"capacity":5,
		"office":"BC Utopia",
		"stage":20,
		"labels":["video","projector"]
	}`
	req, _ := http.NewRequest("POST", "/rooms/create", strings.NewReader(json))

	expected := fmt.Sprintf(`{"id":"%v"}`, stubId)

	response := executeRequest(req, r)

	checkResponseCode(t, http.StatusOK, response.Code)
	require.Equal(t, expected, response.Body.String())
}

func TestCreateRoomWithBadRequest(t *testing.T) {
	// Create a New Server Struct
	r := chi.NewRouter()
	r.Route("/", Make(&logic, logger))

	// Create a New Request
	// typo!
	json := `
	{ 
		"names":"Belyash", 
		"capacity":5,
		"office":"BC Utopia",
		"stage":20,
		"labels":["video","projector"]
	}`
	req, _ := http.NewRequest("POST", "/rooms/create", strings.NewReader(json))

	// Execute Request
	response := executeRequest(req, r)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
	require.Equal(t, "room name can't be empty", response.Body.String())
}

func TestUpdateRoomSuccessfully(t *testing.T) {
	// Create a New Server Struct
	r := chi.NewRouter()
	r.Route("/", Make(&logic, logger))

	// Create a New Request
	json := `
	{
		"id": "123",
		"name":"Belyash",
		"capacity":5,
		"office":"BC Utopia",
		"stage":20,
		"labels":["video","projector"]
	}`
	req, _ := http.NewRequest("POST", "/rooms/update", strings.NewReader(json))
	response := executeRequest(req, r)

	checkResponseCode(t, http.StatusOK, response.Code)
	require.Equal(t, "", response.Body.String())
}

func TestUpdateRoomWithBadRequest(t *testing.T) {
	// Create a New Server Struct
	r := chi.NewRouter()
	r.Route("/", Make(&logic, logger))

	// Create a New Request
	// typo!
	json := `
	{ 
		"names":"Belyash", 
		"capacity":5,
		"office":"BC Utopia",
		"stage":20,
		"labels":["video","projector"]
	}`
	req, _ := http.NewRequest("POST", "/rooms/update", strings.NewReader(json))

	// Execute Request
	response := executeRequest(req, r)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
	require.Equal(t, "room id can't be empty", response.Body.String())
}

func TestDeleteRoomSuccessfully(t *testing.T) {
	// Create a New Server Struct
	r := chi.NewRouter()
	r.Route("/", Make(&logic, logger))

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/rooms/%v", stubId), nil)

	// Execute Request
	response := executeRequest(req, r)

	checkResponseCode(t, http.StatusOK, response.Code)
	require.Equal(t, "", response.Body.String())
}

type logicStub struct{}

var stubId = uuid.New()

func (logicStub) Create(ctx context.Context, room *models.NewRoomInfo) (uuid.UUID, error) {
	return stubId, nil
}

func (logicStub) Update(ctx context.Context, room *models.RoomInfo) error {
	return nil
}

func (logicStub) List(ctx context.Context) ([]models.RoomInfo, error) {
	list := []models.RoomInfo{{Id: stubId.String(), Name: "Belyash", Capacity: 5, Office: "BC Utopia", Stage: 20, Labels: []string{"video", "projector"}}}
	return list, nil
}

func (logicStub) Delete(ctx context.Context, id *uuid.UUID) error {
	return nil
}
