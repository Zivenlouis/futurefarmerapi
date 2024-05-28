package models

import "time"

type SensorData struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	Ph          float64   `gorm:"type:decimal(18,2)" json:"ph"`
	Tds         float64   `gorm:"type:decimal(18,2)" json:"tds"`
	Temperature float64   `gorm:"type:decimal(18,2)" json:"temperature"`
	Humidity    float64   `gorm:"type:decimal(18,2)" json:"humidity"`
	CreatedAt   time.Time `json:"created_at"`
	RelayStatus []RelayStatus `gorm:"foreignKey:SensorId" json:"relay_status"` 
}

type RelayStatus struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	SensorId   int64     `gorm:"not null;index" json:"sensor_id"` // Corrected the type and added json annotation
	Type       string    `gorm:"type:varchar(100);not null" json:"type"`
	Status     int64     `gorm:"type:int;not null" json:"status"` // Ensured the type is consistent and added json annotation
	UpdatedAt  time.Time `json:"updated_at"` // Corrected the field name to follow Go conventions
}

func (RelayStatus) TableName() string {
    return "relay_statuses"
}

func (SensorData) TableName() string {
    return "sensor_data"
}