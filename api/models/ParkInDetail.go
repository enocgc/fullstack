package models

import (
	"time"
)

type ParkInDetail struct {
	ID             uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Horario        string    `gorm:"size:255;not null;" json:"horario"`
	NombreParqueo  string    `gorm:"size:100;not null;" json:"nombreParqueo"`
	Detalle        string    `gorm:"size:300;not null;" json:"detalle"`
	Lat            string    `gorm:"size:100;not null;" json:"lat"`
	Long           string    `gorm:"size:100;not null;" json:"long"`
	Phone          string    `gorm:"size:100;not null;" json:"phone"`
	SitioWeb       string    `gorm:"size:100;not null;" json:"sitioWeb"`
	Fk_ParkInAdmin uint32    `gorm:"not null" json:"fkparkinadmin"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
