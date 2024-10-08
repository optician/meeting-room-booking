package dbPool

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type DbPool interface {
	GetPool() *pgxpool.Pool
	Close()
}

type DBConfig struct {
	Url string `koanf:"url"`
}

// wrapper is for simplification of initialization during wiring
type poolImpl struct {
	Pool *pgxpool.Pool
}

func NewDBPool(config *DBConfig, logger *zap.SugaredLogger) (DbPool, error) {
	dbpool, err := pgxpool.New(context.Background(), config.Url)

	c := context.Background()
	if err := dbpool.Ping(c); err != nil {
		logger.Errorf("db ping failed, ctx: %v, error: %v", c, err)
	} else {
		logger.Info("db ping succeded, ctx: %v", c)
	}

	if err != nil {
		logger.Errorf("cannot create a DB pool, %v", err)
		return nil, err
	} else {
		var pool DbPool = poolImpl{
			Pool: dbpool,
		}
		return pool, nil
	}
}

func (dbPool poolImpl) GetPool() *pgxpool.Pool {
	return dbPool.Pool
}

func (dbPool poolImpl) Close() {
	dbPool.Pool.Close()
}
