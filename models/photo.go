package models

import "github.com/jinzhu/gorm"

type Photo struct{
	gorm.Model
	Url string
	ContentID uint
}

