package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"template/internal/global"
)

const DB_URL = "root:123456@(127.0.0.1:3306)/testdb"

func init() {
	db, err := gorm.Open(mysql.Open(DB_URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	global.Conf.DB = db

}
