package config

import (
	"fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
	"GO14/finalProject/models"
)

var (
	Db    *gorm.DB
	errDB error
)

func init() {
	dsn := "host=localhost user=postgres password=postgres dbname=todos_golang port=5432 sslmode=disable"
	Db, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDB != nil {
		fmt.Println("error:", errDB)
	}

	fmt.Println("db connection established")
	Db.Debug().AutoMigrate(&models.Todos{})
}