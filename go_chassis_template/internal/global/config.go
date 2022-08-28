package global

import "gorm.io/gorm"

var Conf Config

type Config struct {
	DB *gorm.DB
}
