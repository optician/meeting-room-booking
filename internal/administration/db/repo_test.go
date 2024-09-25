package db

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
	"github.com/optician/meeting-room-booking/internal/administration/db/testing"
	"github.com/optician/meeting-room-booking/internal/administration/models"
	"github.com/optician/meeting-room-booking/internal/dbPool"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"go.uber.org/zap"
)

type AdministrationRepositoryTestSuite struct {
	suite.Suite
	pgContainer *postgres.PostgresContainer
	repository  *DB
	ctx         context.Context
	logger      zap.SugaredLogger
}

func (suite *AdministrationRepositoryTestSuite) SetupTest() {
	ctx := context.Background()
	logger := zap.NewExample().Sugar()
	container, err := testHelpers.CreatePostgresContainer(ctx)
	if err != nil {
		logger.Fatalf("cannot setup postgres container in AdministrationRepositoryTestSuite, %v", err)
	}
	config := dbPool.DBConfig{Url: container.ConnectionString}

	dbPool, dbPoolErr := dbPool.NewDBPool(&config, logger)
	if dbPoolErr != nil {
		logger.Fatalf("application terminated: %v", dbPoolErr)
	}

	roomsDB := New(dbPool.GetPool(), logger)

	suite.pgContainer = container.Container
	suite.repository = &roomsDB
	suite.ctx = ctx

	// Migration
	if conn, err := pgx.Connect(ctx, container.ConnectionString); err != nil {
		logger.Fatalf("cannot connect tern to DB, %v", err)
	} else {
		migrator, _ := migrate.NewMigrator(ctx, conn, "public.schema_version")
		migrationsPath := "../../../migrations/" // better to use env instead
		err = migrator.LoadMigrations(os.DirFS(migrationsPath))
		if err != nil {
			logger.Fatalf("Error loading migrations:\n  %v\n", err)
		}
		if len(migrator.Migrations) == 0 {
			logger.Fatalf("No migrations found")
		}
		if err := migrator.Migrate(ctx); err != nil {
			logger.Fatalf("migration failed, %v", err)
		}
	}
}

func (suite *AdministrationRepositoryTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		suite.logger.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *AdministrationRepositoryTestSuite) TestCreateListRooms() {
	newRoom := models.NewRoomInfo{
		Name:     "Cheburechnaya",
		Capacity: 6,
		Office:   "FoodCourt",
		Stage:    -2,
		Labels:   []string{"video", "projector", "whiteboard"},
	}

	id, err := (*suite.repository).Create(suite.ctx, &newRoom)
	require.Nil(suite.T(), err, "Create error")

	expected := models.RoomInfo{
		Id:       id.String(),
		Name:     newRoom.Name,
		Capacity: newRoom.Capacity,
		Office:   newRoom.Office,
		Stage:    newRoom.Stage,
		Labels:   newRoom.Labels,
	}

	rooms, err := (*suite.repository).List(suite.ctx)
	require.Nil(suite.T(), err, "List error")
	require.Equal(suite.T(), []models.RoomInfo{expected}, rooms, "List result")
}

func TestAdministrationRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AdministrationRepositoryTestSuite))
}
