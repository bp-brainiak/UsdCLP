package persistance

import (
	"backend/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Tabler interface {
	TableName() string
}

func GetDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		config.GetValue("backhost"),
		config.GetValue("backuser"),
		config.GetValue("backpass"),
		config.GetValue("backdata"),
		config.GetValue("dbport"),
	)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	return db
}
