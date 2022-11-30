package core

import (
	"fmt"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gotipath.com/gostream/config"
)

type Database struct {
	Connection *gorm.DB
}

// New Database connection
func NewDatabase(config *config.Config) *Database {
	dsn := "clickhouse://default:@localhost:9000/default"
	db, err := gorm.Open(clickhouse.New(clickhouse.Config{
		DSN: dsn,
		// Conn:                      conn,     // initialize with existing database conn
		DisableDatetimePrecision:  false,    // disable datetime64 precision, not supported before clickhouse 20.4
		DontSupportRenameColumn:   false,    // rename column not supported before clickhouse 20.4
		SkipInitializeWithVersion: true,     // smart configure based on used version
		DefaultGranularity:        3,        // 1 granule = 8192 rows
		DefaultCompression:        "LZ4",    // default compression algorithm. LZ4 is lossless
		DefaultIndexType:          "minmax", // index stores extremes of the expression
		DefaultTableEngineOpts:    "ENGINE=MergeTree() ORDER BY tuple()",
	}), &gorm.Config{})

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	return &Database{
		Connection: db,
	}
}
