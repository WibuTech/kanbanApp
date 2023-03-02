package utils

import (
	"kanbanApp/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	USE_SQLITE = true
)

var db *gorm.DB

func ConnectDB() error {
	// connect using gorm pgx
	var dialect gorm.Dialector
	if USE_SQLITE {
		dialect = sqlite.Open("local.db")
	} else {
		dialect = postgres.New(postgres.Config{
			DriverName: "pgx",
			DSN:        os.Getenv("DATABASE_URL"),
		})
	}

	conn, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return err
	}

	conn.AutoMigrate(entity.User{}, entity.Category{}, entity.Task{})
	SetupDBConnection(conn)

	return nil
}

func SetupDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}
