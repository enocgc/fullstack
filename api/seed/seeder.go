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

	for i, _ := range userParkinAdmin {
		err = db.Debug().Model(&models.UserParkinAdmin{}).Create(&userParkinAdmin[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = userParkinAdmin[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
