package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/optician/meeting-room-booking/internal/administration/models"
	"github.com/optician/meeting-room-booking/internal/administration/service"
	"go.uber.org/zap"
)

type Controller struct {
	logger *zap.SugaredLogger
	logic  *service.Logic
}

// mutates router
func Make(logic *service.Logic, logger *zap.SugaredLogger) func(chi.Router) {
	defer logger.Sync()

	controller := Controller{
		logger: logger,
		logic:  logic,
	}
	return controller.routes
}

func (ctrl Controller) routes(r chi.Router) {
	r.Route("/rooms", func(r chi.Router) {
		r.Get("/", ctrl.getRoomsController)
		r.Post("/create", ctrl.createRoomController)
		r.Delete("/{id}", ctrl.deleteRoomController)
		r.Post("/update", ctrl.updateRoomController)
	})
}

func (ctrl Controller) getRoomsController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logicChannel := make(chan []models.RoomInfo)
	go ctrl.getRooms(&logicChannel)

	select {
	case <-ctx.Done():
		return
	case result := <-logicChannel:
		json, err := json.Marshal(result)
		if err == nil {
			w.Header().Add("content-type", "application/json")
			w.Write(json)
		} else {
			ctrl.logger.Errorf("internal error: %v", err)
			w.WriteHeader(500)
		}
	}
}

func (ctrl Controller) getRooms(listener *chan []models.RoomInfo) {
	if list, err := (*ctrl.logic).List(); err != nil {
		ctrl.logger.Errorf("internal error: %v", err)
	} else {
		*listener <- list
	}
}

func (ctrl Controller) createRoomController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logicChannel := make(chan any)

	go func() {
		defer func() { logicChannel <- struct{}{} }()

		if room, err := fromBytesNewRoom(r.Body); err != nil {
			ctrl.logger.Errorf("Bad Request. Invalid NewRoomInfo: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else if json, err := ctrl.createRoom(&room); err != nil {
			ctrl.logger.Errorf("failed to create a new room: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(json))
		}
	}()

	select {
	case <-ctx.Done():
		return
	case <-logicChannel: // no cancelation, looks useless
	}
}

// the idea is to prepare answer and errors
// there is only common logic in the controller method
// maybe it helps to generalize controllers and avoid repetition of a fragile error processing
func (ctrl Controller) createRoom(newRoom *models.NewRoomInfo) (string, error) {
	if id, err := (*ctrl.logic).Create(newRoom); err != nil {
		return "", err
	} else {
		response := CreationResponse{Id: id}
		if json, err := json.Marshal(response); err != nil {
			ctrl.logger.Errorf("room creation failed, %v", err)
			return "", err
		} else {
			return string(json), nil
		}
	}
}

func (ctrl Controller) deleteRoomController(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")

	if id, err := uuid.Parse(strId); err != nil {
		ctrl.logger.Error(`deletion of a room called wit malformed id "%v"`, strId)
		w.WriteHeader(http.StatusBadRequest)
	} else if err := ctrl.deleteRoom(id); err != nil {
		ctrl.logger.Errorf("Deletion of %v room raised error: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (ctrl Controller) deleteRoom(id uuid.UUID) error {
	return (*ctrl.logic).Delete(id)
}

func (ctrl Controller) updateRoomController(w http.ResponseWriter, r *http.Request) {
	if room, err := fromBytesRoom(r.Body); err != nil {
		ctrl.logger.Errorf("Bad Request. Invalid RoomInfo: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else if err := ctrl.updateRoom(&room); err != nil {
		ctrl.logger.Errorf("Update of %v room raised error: %v", room, err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (ctrl Controller) updateRoom(room *models.RoomInfo) error {
	return (*ctrl.logic).Update(room)
}
