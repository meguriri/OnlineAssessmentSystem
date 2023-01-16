package database

import (
	"sync"

	"github.com/meguriri/OnlineAssessmentSystem/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func InitDB() {
	once.Do(func() {
		database, err := gorm.Open(mysql.Open(util.GetConfig().Mysql.Dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		db = database
	})
}
