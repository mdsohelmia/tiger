package core

import (
	"context"

	"github.com/MdSohelMia/tiger/config"
	"github.com/uptrace/go-clickhouse/ch"
)

type Database struct {
	Connection *ch.DB
}

// New Database connection
func NewDatabase(config *config.Config) *Database {
	ctx := context.Background()
	db := ch.Connect(ch.WithDatabase(config.DB.Database))

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	return &Database{
		Connection: db,
	}
}
