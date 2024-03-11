package storage

import (
	"database/sql"

	"github.com/tmazitov/auth_service.git/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Storage struct {
	config *config.StorageConfig
	db     *bun.DB
}

func NewStorage(config *config.StorageConfig) (*Storage, error) {

	var storage *Storage = &Storage{
		config: config,
	}

	if !config.Validate() {
		return nil, ErrInvalidConfig
	}

	if storage.config.URL != "" {
		storage.db = bun.NewDB(
			sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(storage.config.URL))),
			pgdialect.New(),
		)

	} else {
		storage.db = bun.NewDB(
			sql.OpenDB(pgdriver.NewConnector(
				pgdriver.WithAddr(storage.config.Addr),
				pgdriver.WithUser(storage.config.User),
				pgdriver.WithPassword(storage.config.Password),
				pgdriver.WithDatabase(storage.config.Database),
			)),
			pgdialect.New(),
		)
	}

	return storage, nil
}
