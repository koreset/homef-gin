package models

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	Email string
	Name string
	Contents []Content
}

