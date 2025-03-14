package database

import (
	"SysAuthenticate/config"
	"SysAuthenticate/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	configs, err := config.ConfigSet()
	if err != nil {
		return nil, err
	}

	connString := configs.DataBase.StringConnection
	typeDatabase := configs.DataBase.TypeDatabase

	var db *gorm.DB
	switch typeDatabase {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(connString), &gorm.Config{})
	default:
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}
	return db, nil
}
