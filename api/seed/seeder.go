package seed


import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/trydoit9/Go-Test/api/models"
)

var users = []models.User{
	models.User{
		UserId: "Gee",
		UserName: "오지",
		UserMail: "Gee@gmail.com",
		Password: "1",
	},
	models.User{
		UserId: "Wonny",
		UserName: "오지워니",
		UserMail: "Wonny@gmail.com",
		Password: "qwe123qwe123!",
	},
}


func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}