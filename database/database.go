package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	dsn := "host=localhost user=sudarshan512 password=sud12345 dbname=gin_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to db")
}
