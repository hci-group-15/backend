package model

import (
	"context"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnector struct {
	DB *gorm.DB
}

var ctx = context.Background()

func Init() *DBConnector {
	db, err := gorm.Open(sqlite.Open("janishutz-store.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("[DB] Failed to connect to DB with error", err)
	}

	// FIXME: Only used if using sqlite db
	db.Exec("PRAGMA foreign_keys = ON")

	return &DBConnector{
		DB: db,
	}
}
