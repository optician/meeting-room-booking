package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/optician/meeting-room-booking/internal/administration/db"
	"github.com/optician/meeting-room-booking/internal/administration/models"
	"go.uber.org/zap"
)

type Logic interface {
	Create(ctx context.Context, room *models.NewRoomInfo) (uuid.UUID, error)

	Update(ctx context.Context, room *models.RoomInfo) error

	List(ctx context.Context) ([]models.RoomInfo, error)

	Delete(ctx context.Context, id *uuid.UUID) error
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

func (impl impl) Create(ctx context.Context, room *models.NewRoomInfo) (uuid.UUID, error) {
	impl.logger.Infof("recieved a new room %v", *room)
	id, err := (*impl.db).Create(ctx, room) // wrap error
	return id, err
}

func (impl impl) Update(ctx context.Context, room *models.RoomInfo) error {
	impl.logger.Infof("recieved an updated room %v", *room)
	return (*impl.db).Update(ctx, room) // wrap error
}

func (impl impl) List(ctx context.Context) ([]models.RoomInfo, error) {
	list, err := (*impl.db).List(ctx) // wrap error
	return list, err
}

func (impl impl) Delete(ctx context.Context, id *uuid.UUID) error {
	impl.logger.Infof("delete %v room", id)
	return (*impl.db).Delete(ctx, id) // wrap error
}
