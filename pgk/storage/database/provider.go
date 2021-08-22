package database

import (
	"fmt"
	"github.com/tonsV2/todo-go/pgk/configuration"
	"github.com/tonsV2/todo-go/pgk/group"
	"github.com/tonsV2/todo-go/pgk/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func ProvideDatabase(configuration configuration.Configuration) *gorm.DB {
	host := configuration.Postgresql.Host
	port := configuration.Postgresql.Port
	username := configuration.Postgresql.Username
	password := configuration.Postgresql.Password
	name := configuration.Postgresql.DatabaseName

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, username, password, name, port)

	databaseConfig := gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	}

	db, err := gorm.Open(postgres.Open(dsn), &databaseConfig)

	if err != nil {
		log.Fatalf("Failed to connect to database!")
	}

	err = db.AutoMigrate(
		&user.User{},
		&group.Group{},
	)

	if err != nil {
		log.Fatalf("Migration error: %s", err)
	}

	return db
}
