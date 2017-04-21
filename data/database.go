package data

// https://siongui.github.io/2016/01/09/go-sqlite-example-basic-usage/

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/trickierstinky/slack-invite-api/config"
)

func SetupDatabase() {
	db := openDatabase()
	defer closeDatabase(db)
	// Migrate the schema
	db.AutoMigrate(&Invite{})
}

func CreateInvite(invite Invite) Invite {
	db := openDatabase()
	defer closeDatabase(db)

	db.Create(&invite)
	return invite
}

func FetchInvite(id int) Invite {
	db := openDatabase()
	defer closeDatabase(db)

	var invite Invite
	db.Where("ID = ?", id).First(&invite)

	return invite
}

func DeleteInvite(invite Invite) bool {
	db := openDatabase()
	defer closeDatabase(db)

	db.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(&invite)

	return true
}

func openDatabase() *gorm.DB {

	db, err := gorm.Open(config.Env("db_provider"), config.Env("db_connection"))

	if err != nil {
		panic(err)
	}

	return db
}

func closeDatabase(db *gorm.DB) {
	db.Close()
}
