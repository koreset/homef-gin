package models

type Link struct {
	ID       uint `gorm:"primary_key"`
	Url      string
	Title    string
	ImageUrl string
	PostID   uint
}

type Video struct {
	ID          uint `gorm:"primary_key"`
	Url         string
	Value       string
	Description string
	PostID      uint
}

type Image struct {
	ID       uint `gorm:"primary_key"`
	Url      string
	FileName string
	PostID   uint
}

type Post struct {
	ID      uint   `gorm:"primary_key"`
	Title   string
	Body    string `gorm:"type:longtext"`
	Summary string `gorm:"type:varchar(500)"`
	Images  []Image
	Videos  []Video
	Links   []Link
	Type    string
	Created int32
	Updated int32
}
