package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Content struct {
	Id          uint `gorm:"primary_key"`
	ContentType string
	ContentId   uint
}

type Article struct {
	Id      uint `gorm:"primary_key"`
	Body    string
	Content Content `gorm:"polymorphic:Content;auto_preload"`
}

type Video struct {
	Id          uint `gorm:"primary_key"`
	Description string
	Url         string
	Content     Content `gorm:"polymorphic:Content;auto_preload"`
}

func main() {
	db, err := gorm.Open("mysql", "root:wordpass15@tcp(localhost:3306)/newdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Content{}, &Article{}, &Video{})
	db.LogMode(true)

	var article Article
	var content Content
	var vid Video

	db.Model(&article).Related(&content)
	db.Model(&vid).Related(&content)

	video := Video{
		Description: "a new video has just been released",
		Url:         "http://example.com/video1.mp4",
		Content: Content{
			ContentType: "video",
		},
	}

	db.Create(&video)
	db.Save(&video)

	var newVideo Video

	db.Preload("Content").Where("videos.id = ?", 2).First(&newVideo)

	fmt.Println(newVideo)

	post := Article{
		Body: "My body is getting better and better version 3",
		Content: Content{
			ContentType: "article",
		},
	}

	db.Create(&post)
	db.Save(&post)

	//var newpost Article
	//db.Where("articles.id = ?", 6).Preload("Content").First(&newpost)
	//
	//fmt.Println(newpost)
	//
	//fmt.Println(newpost.Content.ContentType)

}
