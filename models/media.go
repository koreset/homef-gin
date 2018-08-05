package models

import "github.com/qor/media/media_library"

type Media struct {
	ID               uint `gorm:"primary_key"`
	EntityId         uint
	FieldImageFid    uint
	FieldImageWidth  uint
	FieldImageHeight uint
	Filename         string
	Uri              string
}


type MediaLibrary struct {
	Title string
	media_library.MediaLibrary
}