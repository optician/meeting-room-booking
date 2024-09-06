package administration

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Controller struct {
	logger *zap.SugaredLogger
}

// mutates router
func Make(r chi.Router) {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	controller := Controller{
		logger: logger,
	}
	controller.routes(r)
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
	logicChannel := make(chan []RoomInfo)
	go getRooms(&logicChannel)

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

func getRooms(listener *chan []RoomInfo) {
	*listener <- []RoomInfo{RoomInfo{Id: "123", Name: "Belyash", Capacity: 5, Office: "BC Utopia", Stage: 20, Labels: []string{"video", "projector"}}}
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
		} else if err := ctrl.createRoom(&room); err != nil {
			ctrl.logger.Errorf("failed to create a new room: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}()

	select {
	case <-ctx.Done():
		return
	case <-logicChannel: // no cancelation, looks useless
	}
}

func (ctrl Controller) createRoom(newRoom *NewRoomInfo) error {
	ctrl.logger.Infof("recieved a new room %v", *newRoom)
	return nil
}

func (ctrl Controller) deleteRoomController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		ctrl.logger.Error("deletion of a room called without id")
		w.WriteHeader(http.StatusBadRequest)
	} else if err := deleteRoom(id); err != nil {
		ctrl.logger.Errorf("Deletion of %v room raised error: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func deleteRoom(_ string) error {
	return nil
}

func (ctrl Controller) updateRoomController(w http.ResponseWriter, r *http.Request) {
	if room, err := fromBytesRoom(r.Body); err != nil {
		ctrl.logger.Errorf("Bad Request. Invalid RoomInfo: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else if err := updateRoom(&room); err != nil {
		ctrl.logger.Errorf("Update of %v room raised error: %v", room, err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func updateRoom(_ *RoomInfo) error {
	return nil
}
