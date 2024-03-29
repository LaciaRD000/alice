package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Open(dsn string) (err error) {
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.AutoMigrate(&Ticket{}, &Verify{}, &AntiSpam{}, &Welcome{}, &Reaction{}, &Shop{}, &StatusPanel{}, &Leave{}, &LevelConfig{}, &UserLevel{})
}

/*
func FindAll(query string, args ...string) (tickets []Ticket, err error) {
	err = db.Where(query, args).Find(&tickets).Error
	return tickets, err
}
*/
