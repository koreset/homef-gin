package models

import "time"

type Content struct {
	Nid         uint `gorm:"primary_key"`
	Type        string
	Title       string `gorm:not null`
	BodyValue   string `gorm:"type:longtext"`
	BodySummary string `gorm:"type:longtext"`
	Created     uint32
	Changed     uint32
	Media       []Media `gorm:"foreignKey:EntityId"`
	DeletedAt   *time.Time
}
