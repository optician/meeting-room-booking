package httpapi

import "github.com/google/uuid"

type CreationResponse struct {
	Id uuid.UUID `json:"id"`
}
