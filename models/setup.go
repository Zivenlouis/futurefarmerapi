package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3307)/futurefarmerapi?parseTime=true"))
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(&User{}, &SensorData{}, &RelayStatus{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&SensorData{})
	db.AutoMigrate(&RelayStatus{})
	DB = db
	CreateSensorData()
	createRelayStatus()
}

func CreateSensorData() {
	sensorData := SensorData{
		Ph:          7.5,
		Tds:         500.0,
		Temperature: 25.0,
		Humidity:    60.0,
		CreatedAt:   time.Now(),
	}

	result := DB.Create(&sensorData)
	if result.Error != nil {
		log.Fatalf("failed to create sensor data: %v", result.Error)
	} else {
		fmt.Printf("SensorData created: %+v\n", sensorData)
	}
}

func createRelayStatus() {
	relayStatus := RelayStatus{
		SensorId:   1,
		Type: "ph",
		Duration: 3,
		Status: 0,
		UpdatedAt: time.Now(),
	}
	DB.Create(&relayStatus)

	relayStatus2 := RelayStatus{
		SensorId:   1,
		Type: "tds",
		Duration: 5,
		Status: 0,
		UpdatedAt: time.Now(),
	}
	DB.Create(&relayStatus2)

	relayStatus3 := RelayStatus{
		SensorId:   1,
		Type: "temperature",
		Duration: 10,
		Status: 0,
		UpdatedAt: time.Now(),
	}
	DB.Create(&relayStatus3)

	relayStatus4 := RelayStatus{
		SensorId:   1,
		Type: "humidity",
		Duration: 15,
		Status: 0,
		UpdatedAt: time.Now(),
	}
	DB.Create(&relayStatus4)
}
