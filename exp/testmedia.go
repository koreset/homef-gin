package main

import (
	"github.com/qor/media/media_library"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"github.com/qor/media"
)

type Post struct {
	gorm.Model
	Name string
	Files []Media
}

type Media struct {
	gorm.Model
	File media_library.MediaLibraryStorage `gorm:"type:longtext" sql:"size:4294967295;" media_library:"url:/mytestfiles/{{class}}/{{primary_key}}/{{column}}.{{extension}}"`
	PostID uint
}

func (Media) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":  {Width: 320, Height: 320},
		"middle": {Width: 640, Height: 640},
		"big":    {Width: 1280, Height: 1280},
		"preview": {Width:200, Height:200},
	}
}


func main() {
	localConnection := "root:wordpass15@tcp(localhost:3306)/newtest?charset=utf8&parseTime=True&loc=Local"
	workingDirectory := "/Users/jome/projects/homef/files/"

	db, _ := gorm.Open("mysql", localConnection)

	var post Post
	var newmedia Media
	db.LogMode(true)
	db.AutoMigrate(&Post{}, &Media{})
	db.Model(&post).Related(&newmedia)
	media.RegisterCallbacks(db)

	file, err := os.Open(workingDirectory + "biosafety-act.png")
	if err != nil {
		panic(err)
	}
	newmedia = Media{}
	defer file.Close()

	newmedia.File.Sizes = newmedia.GetSizes()
	newmedia.File.Scan(file)

	post = Post{Name: "TestPost 4"}
	post.Files = append(post.Files, newmedia)

	db.Save(&post)


}
