package jobs

import (
	"gcode/orm/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var (
	uname                    = "admin"
	fname                    = "admin"
	lname                    = "admin"
	nname                    = "admin"
	description              = "This is the admin user ever!"
	location                 = "His house, maybe? Wouldn't know"
	password                 = "admin"
	firstUser   *models.User = &models.User{
		Email:       "mkinsz@hotmail.com",
		Name:        &uname,
		FirstName:   &fname,
		LastName:    &lname,
		NickName:    &nname,
		Description: &description,
		Location:    &location,
		Password:    &password,
	}
)

// SeedUsers inserts the first users
var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstUser).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstUser).Error
	},
}
