package seed

import (
	"log"

	"github.com/enocgc/fullstack/api/models"
	"github.com/jinzhu/gorm"
)

var userParkinAdmin = []models.UserParkinAdmin{
	models.UserParkinAdmin{
		Username:  "Steven victor",
		Email:     "steven@gmail.com",
		Phone:     "83873481",
		Password:  "password",
		TipoLogin: "Facebook",
	},
	models.UserParkinAdmin{
		Username:  "Martin Luther",
		Email:     "luther@gmail.com",
		Phone:     "84059516",
		Password:  "password",
		TipoLogin: "Google",
	},
}

var userParkinClient = []models.UserParkinClient{
	models.UserParkinClient{
		Name:         "Steven ",
		LastName:     "victor",
		Email:        "steven@gmail.com",
		Phone:        "83873481",
		Password:     "password",
		TipoRegistro: "Facebook",
		Token:        "12345",
	},
	models.UserParkinClient{
		Name:         "Martin Luther",
		LastName:     "King",
		Email:        "luther@gmail.com",
		Phone:        "84059516",
		Password:     "password",
		TipoRegistro: "Google",
		Token:        "12345",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}
var parkInDetail = []models.ParkInDetail{
	models.ParkInDetail{
		Horario:       "lunea a viernes de 8:00 am a 10:00 pm",
		NombreParqueo: "Parque Prkin V1 1",
		Detalle:       "Este es el primer parqueo",
		Lat:           "1",
		Long:          "2",
		Phone:         "83873481",
		SitioWeb:      "parkIn.com",
	},
	models.ParkInDetail{
		Horario:       "lunea a viernes de 8:00 am a 10:00 pm",
		NombreParqueo: "Parque Prkin V1 2",
		Detalle:       "Este es el primer parqueo2",
		Lat:           "1",
		Long:          "2",
		Phone:         "83873481",
		SitioWeb:      "parkIn2.com",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.UserParkinAdmin{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.UserParkinAdmin{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "user_parkin_admins(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	// ParkInDetail
	// err = db.Debug().DropTableIfExists(&models.ParkInDetail{}, &models.UserParkinAdmin{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot drop table: %v", err)
	// }
	err = db.Debug().AutoMigrate(&models.UserParkinAdmin{}, &models.ParkInDetail{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.ParkInDetail{}).AddForeignKey("fk_park_in_admin", "user_parkin_admins(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range userParkinAdmin {
		err = db.Debug().Model(&models.UserParkinAdmin{}).Create(&userParkinAdmin[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = userParkinAdmin[i].ID
		parkInDetail[i].Fk_ParkInAdmin = userParkinAdmin[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
		err = db.Debug().Model(&models.ParkInDetail{}).Create(&parkInDetail[i]).Error
		if err != nil {
			log.Fatalf("cannot seed parkInDetail table: %v", err)
		}
	}

	// crear clientes db
	err = db.Debug().AutoMigrate(&models.UserParkinClient{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for j, _ := range userParkinClient {
		// crear clientes db
		err = db.Debug().Model(&models.UserParkinClient{}).Create(&userParkinClient[j]).Error
		if err != nil {
			log.Fatalf("cannot seed user_parkin_client table: %v", err)
		}
	}

}
