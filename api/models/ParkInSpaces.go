package models

import (
	"time"
)

type ParkInSpaces struct {
	ID                 uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CantidadTotal      int       `gorm:"size:255;not null;" json:"cantidadtotal"`
	CantidadOcupados   int       `gorm:"size:255;not null;" json:"cantidadocupados"`
	CantidadReservados int       `gorm:"size:255;not null;" json:"cantidadreservados"`
	Fk_ParkInAdmin     uint32    `gorm:"not null" json:"fkparkinadmin"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
