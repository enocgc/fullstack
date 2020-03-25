package models

import (
	"time"
)

type ParkInProcess struct {
	ID              uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Cantidad        int       `gorm:"size:255;not null;" json:"cantidad"`
	Status          int       `gorm:"size:100;not null;" json:"status"`
	Fk_ParkInReserv uint32    `gorm:"not null" json:"fkparkinreserv"`
	Fk_ParkSpaces   uint32    `gorm:"not null" json:"fkparkinspaces"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
