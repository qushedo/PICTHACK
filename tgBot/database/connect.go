package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"tgBot/entities"
	"time"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=database port=5432 sslmode=disable TimeZone=UTC", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	} else {
		log.Println("Успешно подключились к базе данных")
	}

	err = database.AutoMigrate(
		&entities.User{},
		&entities.Course{},
	)
	if err != nil {
		log.Panic(err)
	}

	return database
}
