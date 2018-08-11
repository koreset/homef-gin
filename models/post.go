package models

import (
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/qor/media/media_library"
	"github.com/qor/media"
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
	File   media_library.MediaLibraryStorage `gorm:"type:longtext" sql:"size:4294967295;" media_library:"url:/content/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./public"`
	PostID uint
}

func (Image) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":           {Width: 320, Height: 320},
		"middle":          {Width: 640, Height: 640},
		"big":             {Width: 1280, Height: 1280},
		"article_preview": {Width: 390, Height: 300},
		"preview":         {Width: 200, Height: 200},
	}
}

type Document struct {
	gorm.Model
	File   oss.OSS `gorm:"type:longtext" sql:"size:4294967295;" media_library:"url:/content/publications/{{basename}}.{{extension}};path:./public"`
	PostID uint
}

type Post struct {
	ID        uint   `gorm:"primary_key"`
	Title     string
	Body      string `gorm:"type:longtext"`
	Summary   string `gorm:"type:longtext"`
	Images    []Image
	Documents []Document
	Videos    []Video
	Links     []Link
	Type      string
	Created   int32
	Updated   int32
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
