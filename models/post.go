package models

import (
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
)

type Link struct {
	gorm.Model
	Url      string
	Title    string
	ImageUrl string
	PostID   uint
}

type Video struct {
	gorm.Model
	Url         string
	Value       string
	Description string
	PostID      uint
}

//type Image struct {
//	ID       uint `gorm:"primary_key"`
//	Url      string
//	FileName string
//	PostID   uint
//}

type Image struct {
	gorm.Model
	ImageFile oss.OSS `sql:"size:4294967295;" media_library:"url:/content/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./public"`
	//ImageFile oss.OSS
	PostID    uint
}

type Post struct {
	ID      uint   `gorm:"primary_key"`
	Title   string
	Body    string `gorm:"type:longtext"`
	Summary string `gorm:"type:longtext"`
	Images  []Image
	Videos  []Video
	Links   []Link
	Type    string
	Created int32
	Updated int32
}

func (p *Post) BeforeCreate() (err error) {
	fmt.Println(p.Body)
	if p.Created == 0 {
		p.Created = int32(time.Now().Unix())
	}

	if p.Updated == 0 {
		p.Updated = int32(time.Now().Unix())
	}

	return nil
}
