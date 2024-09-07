package httpapi

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSerializeCreationResponse(t *testing.T) {
	data := CreationResponse{Id: uuid.New()}
	expected := fmt.Sprintf(`{"id":"%v"}`, data.Id)
	actual, _ := json.Marshal(data)
	require.Equal(t, expected, string(actual))
}
