package models

type Media struct {
	ID               uint `gorm:"primary_key"`
	EntityId         uint
	FieldImageFid    uint
	FieldImageWidth  uint
	FieldImageHeight uint
	Filename         string
	Uri              string
}
