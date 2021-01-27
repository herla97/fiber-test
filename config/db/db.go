package db

import (
	"fiapi/config"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connection gorm connector
var Connection *gorm.DB
var err error

// Connect connect to db
func Connect() {
	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Env("DB_HOST"),
		config.Env("DB_PORT"),
		config.Env("DB_USERNAME"),
		config.Env("DB_NAME"),
		config.Env("DB_PASSWORD"),
	)

	Connection, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if config.Env("APP_ENV") == "dev" {
		Connection.Logger.LogMode(logger.Info)
	}

	sqlDB, err := Connection.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Connection Opened to Database")
}

// Migrate migrates all the database tables
func Migrate(tables ...interface{}) error {
	return Connection.AutoMigrate(tables...)
}
