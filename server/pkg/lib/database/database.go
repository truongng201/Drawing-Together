package database

import (
	"server/pkg/config"
	"server/pkg/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {}

var DB Database

func genDSN() string{
	dsn := "postgres://%s:%s@%s:%s/%s?sslmode=disable"
	dbConfig := config.Con.Database
	return fmt.Sprintf(
		dsn, 
		dbConfig.POSTGRES_USER, 
		dbConfig.POSTGRES_PASSWORD, 
		dbConfig.POSTGRES_HOST,
		dbConfig.POSTGRES_PORT, 
		dbConfig.POSTGRES_DB,
	)
}

func (database Database) Connect() *gorm.DB {
	dsn := genDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database")
		return nil
	}

	db.AutoMigrate(&model.User{})
	fmt.Println("Database connected")
	return db
}
