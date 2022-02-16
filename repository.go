package lib_db

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}
