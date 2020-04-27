package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ParkInDetail struct {
	ID             uint32          `gorm:"primary_key;auto_increment" json:"id"`
	Horario        string          `gorm:"size:255;not null;" json:"horario"`
	NombreParqueo  string          `gorm:"size:100;not null;" json:"nombreParqueo"`
	Detalle        string          `gorm:"size:300;not null;" json:"detalle"`
	Lat            string          `gorm:"size:100;not null;" json:"lat"`
	Long           string          `gorm:"size:100;not null;" json:"long"`
	Phone          string          `gorm:"size:100;not null;" json:"phone"`
	SitioWeb       string          `gorm:"size:100;not null;" json:"sitioWeb"`
	ParkInAdmin    UserParkinAdmin `json:"parkinAdmin"`
	Fk_ParkInAdmin uint32          `gorm:"not null" json:"fk_park_in_admin"`
	CreatedAt      time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *ParkInDetail) PrepareDetail() {
	u.ID = 0
	u.Horario = html.EscapeString(strings.TrimSpace(u.Horario))
	u.NombreParqueo = html.EscapeString(strings.TrimSpace(u.NombreParqueo))
	u.Detalle = html.EscapeString(strings.TrimSpace(u.Detalle))
	u.Phone = html.EscapeString(strings.TrimSpace(u.Phone))
	u.SitioWeb = html.EscapeString(strings.TrimSpace(u.SitioWeb))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *ParkInDetail) ValidateDetail() error {

	if u.Horario == "" {
		return errors.New("Required horario")
	}
	if u.NombreParqueo == "" {
		return errors.New("Required nombre parqueo")
	}
	if u.Phone == "" {
		return errors.New("Required phone")
	}
	if u.Detalle == "" {
		return errors.New("Required Detalle")
	}
	if u.SitioWeb == "" {
		return errors.New("Required Sitio web")
	}

	return nil

}

func (p *ParkInDetail) SaveParkinDetail(db *gorm.DB) (*ParkInDetail, error) {
	var err error
	err = db.Debug().Model(&ParkInDetail{}).Create(&p).Error
	if err != nil {
		return &ParkInDetail{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&UserParkinAdmin{}).Where("id = ?", p.Fk_ParkInAdmin).Take(&p.ParkInAdmin).Error
		if err != nil {
			return &ParkInDetail{}, err
		}
	}
	return p, nil
}

func (p *ParkInDetail) FindAllParkin(db *gorm.DB) (*[]ParkInDetail, error) {
	var err error
	parkin := []ParkInDetail{}
	err = db.Debug().Model(&ParkInDetail{}).Limit(100).Find(&parkin).Error
	if err != nil {
		return &[]ParkInDetail{}, err
	}
	if len(parkin) > 0 {
		for i, _ := range parkin {
			err := db.Debug().Model(&UserParkinAdmin{}).Where("id = ?", parkin[i].Fk_ParkInAdmin).Take(&parkin[i].ParkInAdmin).Error
			if err != nil {
				return &[]ParkInDetail{}, err
			}
		}
	}
	return &parkin, nil
}
