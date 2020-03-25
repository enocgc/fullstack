package models

import (
	"time"
)

type ParkInReserv struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Token        string    `gorm:"size:255;not null;" json:"token"`
	Expiracion   string    `gorm:"size:100;not null;" json:"expiracion"`
	Fk_IdCliente uint32    `gorm:"not null" json:"fkidcliente"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
