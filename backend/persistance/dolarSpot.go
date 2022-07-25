package persistance

import (
	"backend/config"
	"backend/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Tabler interface {
	TableName() string
}

func PersistDolarInfo(info model.SpotInfo) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		config.GetValue("backhost"),
		config.GetValue("backuser"),
		config.GetValue("backpass"),
		config.GetValue("backdata"),
		config.GetValue("dbport"),
	)
	log.Println(dsn)
	//dbport
	//dsn := "host=" + config.GetValue("backhost") + " user=" + +" password=" + config - config.GetValue("backpass")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	result := db.Create(&info)
	log.Println(result)
}
