package service

import (
	"github.com/google/uuid"
	"github.com/optician/meeting-room-booking/internal/administration/db"
	"github.com/optician/meeting-room-booking/internal/administration/models"
	"go.uber.org/zap"
)

type Logic interface {
	Create(room *models.NewRoomInfo) (uuid.UUID, error)

	Update(room *models.RoomInfo) error

	List() ([]models.RoomInfo, error)

	Delete(id uuid.UUID) error
}

type impl struct {
	logger *zap.SugaredLogger
	db     *db.DB
}

func Make(db *db.DB, logger *zap.SugaredLogger) Logic {
	defer logger.Sync()

	return impl{
		logger: logger,
		db:     db,
	}
}

func (impl impl) Create(room *models.NewRoomInfo) (uuid.UUID, error) {
	impl.logger.Infof("recieved a new room %v", *room)
	id, err := (*impl.db).Create(*room) // wrap error
	return id, err
}

func (impl impl) Update(room *models.RoomInfo) error {
	impl.logger.Infof("recieved an updated room %v", *room)
	return (*impl.db).Update(*room) // wrap error
}

func (impl impl) List() ([]models.RoomInfo, error) {
	list, err := (*impl.db).List() // wrap error
	return list, err
}

func (impl impl) Delete(id uuid.UUID) error {
	impl.logger.Infof("delete %v room", id)
	return (*impl.db).Delete(id) // wrap error
}
