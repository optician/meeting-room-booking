package db

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/optician/meeting-room-booking/internal/administration/models"
	"go.uber.org/zap"
)

type DB interface {
	List(context.Context) ([]models.RoomInfo, error)
	Update(context.Context, *models.RoomInfo) error
	Create(context.Context, *models.NewRoomInfo) (uuid.UUID, error)
	Delete(context.Context, *uuid.UUID) error
}

type impl struct {
	logger *zap.SugaredLogger
	dbpool *pgxpool.Pool
}

func New(dbPool *pgxpool.Pool, logger *zap.SugaredLogger) DB {
	return &impl{
		logger: logger,
		dbpool: dbPool,
	}
}

// slice can't be nil if error is nil
func (impl *impl) List(ctx context.Context) ([]models.RoomInfo, error) {
	list := make([]models.RoomInfo, 0)
	query := "select id, name, capacity, office, stage, labels from meeting_rooms"
	err := pgxscan.Select(ctx, impl.dbpool, &list, query)
	return list, err // wrap error
}

func (impl *impl) Update(ctx context.Context, room *models.RoomInfo) error {
	query := `update meeting_rooms 
				set 
					name = @name,
					capacity = @capacity,
					office = @office,
					stage = @stage,
					labels = @labels
				where id = @id`
	args := pgx.NamedArgs{
		"id":       room.Id,
		"name":     room.Name,
		"capacity": room.Capacity,
		"office":   room.Office,
		"stage":    room.Stage,
		"labels":   room.Labels,
	}
	_, err := impl.dbpool.Exec(ctx, query, args)
	return err // wrap error, check existance error
}

func (impl *impl) Create(ctx context.Context, room *models.NewRoomInfo) (uuid.UUID, error) {
	id := uuid.New()
	query := `insert into meeting_rooms 
				(
					id, 
					name,  
					capacity, 
					office,
					stage, 
					labels 
				)
				values (
					@id,
					@name,
					@capacity,
					@office,
					@stage,
					@labels
				)
				`
	args := pgx.NamedArgs{
		"id":       id,
		"name":     room.Name,
		"capacity": room.Capacity,
		"office":   room.Office,
		"stage":    room.Stage,
		"labels":   room.Labels,
	}
	_, err := impl.dbpool.Exec(ctx, query, args)
	return id, err // wrap error, check constraints violations
}

func (impl *impl) Delete(ctx context.Context, id *uuid.UUID) error {
	query := "delete from meeting_rooms where id = @id"
	args := pgx.NamedArgs{"id": id}
	_, err := impl.dbpool.Exec(ctx, query, args)
	return err // wrap error
}
