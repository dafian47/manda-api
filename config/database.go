package config

import (
	"github.com/dafian47/manda-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func InitDB(databaseUrl string) *gorm.DB {

	db, err := gorm.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	// Enable Log Mode if you want to enable Log on Database Query ( Gorm )
	// And disable Log Mode if you want to deploy to Production
	if IsDevelopment {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}

	db.AutoMigrate(
		// Table Master
		&model.MandaMarriageStatus{},
		&model.MandaMajor{},
		&model.MandaRole{},
		&model.MandaWork{},
		&model.MandaUserStatus{},
		&model.MandaThreadStatus{},
		// Table Primary
		&model.MandaAuth{},
		&model.MandaUser{},
		&model.MandaChannel{},
		&model.MandaThread{},
		//&model.MandaComment{},
		//&model.MandaSocialAccount{},
	)

	return db
}
