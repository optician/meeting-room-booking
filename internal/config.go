package internal

import "github.com/optician/meeting-room-booking/internal/dbPool"

type Config struct {
	DB dbPool.DBConfig `koanf:"db"`
}
