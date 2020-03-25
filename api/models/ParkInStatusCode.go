package models

import (
	"time"
)

type ParkInStatusCode struct {
	ID               uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Ingreso          string    `gorm:"size:255;not null;" json:"ingreso"`
	Salida           string    `gorm:"size:100;not null;" json:"salida"`
	Monto            string    `gorm:"size:100;not null;" json:"monto"`
	Fk_ParkInProcess uint32    `gorm:"not null" json:"fkparkinprocess"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
