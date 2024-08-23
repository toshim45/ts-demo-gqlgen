package tools

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbOnce sync.Once
var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		log.Println("[db] init connection")
		dbOnce.Do(func() {
			var err error
			dsn := "root:root@tcp(127.0.0.1:3306)/trial?parseTime=True"
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Fatalln("[db] connect db FAIL", err)
			}
		})
	}
	return db
}
