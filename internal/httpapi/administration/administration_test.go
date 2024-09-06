package administration

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

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
    r.Route("/",Make)

    req, _ := http.NewRequest("GET", "/rooms", nil)

    // Execute Request
    response := executeRequest(req, r)

    // Check the response code
    checkResponseCode(t, http.StatusOK, response.Code)
	expected := `[{"id":"123","name":"Belyash","capacity":5,"office":"BC Utopia","stage":20,"labels":["video","projector"]}]`
    // We can use testify/require to assert values, as it is more convenient
    require.Equal(t, expected, response.Body.String())
}

func TestCreateRoomSuccessfully(t *testing.T) {
    // Create a New Server Struct
	r := chi.NewRouter()
    r.Route("/",Make)

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

    // Execute Request
    response := executeRequest(req, r)

    // Check the response code
    checkResponseCode(t, http.StatusOK, response.Code)

    // We can use testify/require to assert values, as it is more convenient
    require.Equal(t, "", response.Body.String())
}

func TestCreateRoomWithBadRequest(t *testing.T) {
    // Create a New Server Struct
	r := chi.NewRouter()
    r.Route("/",Make)

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

    // Check the response code
    checkResponseCode(t, http.StatusBadRequest, response.Code)

    // We can use testify/require to assert values, as it is more convenient
    require.Equal(t, "room name can't be empty", response.Body.String())
}

func TestUpdateRoomSuccessfully(t *testing.T) {
    // Create a New Server Struct
	r := chi.NewRouter()
    r.Route("/",Make)

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
    req, _ := http.NewRequest("POST", "/rooms/create", strings.NewReader(json))

    // Execute Request
    response := executeRequest(req, r)

    // Check the response code
    checkResponseCode(t, http.StatusOK, response.Code)

    // We can use testify/require to assert values, as it is more convenient
    require.Equal(t, "", response.Body.String())
}

func TestUpdateRoomWithBadRequest(t *testing.T) {
    // Create a New Server Struct
	r := chi.NewRouter()
    r.Route("/",Make)

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

    // Check the response code
    checkResponseCode(t, http.StatusBadRequest, response.Code)

    // We can use testify/require to assert values, as it is more convenient
    require.Equal(t, "room id can't be empty", response.Body.String())
}

func TestDeleteRoomSuccessfully(t *testing.T) {
    // Create a New Server Struct
	r := chi.NewRouter()
    r.Route("/",Make)

    req, _ := http.NewRequest("DELETE", "/rooms/541", nil)

    // Execute Request
    response := executeRequest(req, r)

    // Check the response code
    checkResponseCode(t, http.StatusOK, response.Code)

    // We can use testify/require to assert values, as it is more convenient
    require.Equal(t, "", response.Body.String())
}
