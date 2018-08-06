package main

import (
	"os"
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
	"github.com/qor/media"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/qor/oss/s3"
)

type Product struct {
	gorm.Model
	Images []Image
}

type Image struct {
	gorm.Model
	ImageFile oss.OSS `sql:"size:4294967295;" media_library:"url:/content/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./public"`
	//ImageFile oss.OSS
	ProductID uint
}

func main() {
	localConnection := "root:wordpass15@tcp(localhost:3306)/stores?charset=utf8&parseTime=True&loc=Local"
	newDB, dbError := gorm.Open("mysql", localConnection)
	if dbError != nil {
		panic(dbError)
	}

	newDB.AutoMigrate(&Product{}, &Image{})
	media.RegisterCallbacks(newDB)

	workingDir := "/Users/jome/projects/homef/files/"
	myfile, _ := os.Open(workingDir + "cakexlogo.png")
	var post Product
	var images []Image

	image := Image{
		ImageFile: oss.OSS{
			media.Base{
				Reader: myfile,
			},
		},
	}

	images = append(images, image)
	post = Product{
		Images: images,
	}

	newDB.Save(&post)

	storage := s3.New(s3.Config{AccessID: "access_id", AccessKey: "access_key", Region: "region", Bucket: "bucket", Endpoint: "cdn.getqor.com", ACL: awss3.BucketCannedACLPublicRead})
	//wd, _ := os.Getwd()
	//workingDir := wd + "/public/content"
	//workingDir := "/Users/jome/projects/homef/files"
	//storage := filesystem.New(workingDir)
	//
	//file, err := os.Open(workingDir + "/abnlogo.png")
	//if err != nil {
	//	panic(err)
	//}
	//storage.Put("/mycontent/sample.png", file)
	//fmt.Println(workingDir)

}
