package persistance

import (
	"backend/model/spot"
	"log"
)

func PersistDolarInfo(info spot.SpotInfo) {

	db := GetDatabase()
	defer CloseConnection(db)
	result := db.Create(&info)
	log.Println(result)

}
