package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/koreset/homefnew/app/models"
	"strings"
	"github.com/gosimple/slug"
)

var dbError error
var newDB *gorm.DB

func transformString(file string) string {
	parts := strings.Split(file, ".")
	extension := parts[len(parts)-1]
	filename := strings.TrimSuffix(file, "."+extension)
	modFileName := slug.MakeLang(filename, "en") + "." + extension

	return modFileName

}

func main() {
	//awsConnection := "homef:wordpass15@tcp(rds-mysql-homef.cb44dbuhyviz.eu-west-2.rds.amazonaws.com:3306)/homef?charset=utf8&parseTime=True&loc=Local"
	localConnection := "root:wordpass15@tcp(localhost:3306)/homef?charset=utf8&parseTime=True&loc=Local"
	//directory := "/Users/jome/go/src/github.com/koreset/homef-gin/public/content/"

	newDB, dbError = gorm.Open("mysql", localConnection)

	if dbError != nil {
		panic(dbError.Error())
	}
	defer newDB.Close()

	var images []models.Image
	newDB.Raw("select * from images").Find(&images)

	for _, v := range images {
		originalPath := v.Url
		orig_fileName := v.FileName

		original_parts := strings.Split(originalPath, "/")
		o_filename := original_parts[len(original_parts)-1]

		newfileName := transformString(o_filename)
		newFilePath := strings.Replace(originalPath, o_filename, newfileName, -1)

		fmt.Println(originalPath)
		fmt.Println(newFilePath)
		fmt.Println(orig_fileName)
		fmt.Println()

		fmt.Println()
		v.Url = newFilePath
		v.FileName = newfileName

		newDB.Save(&v)

		//file, err := os.Open(directory + v.Url)
		//if err != nil {
		//	fmt.Println(err.Error())
		//	continue
		//}
		//
		//info, _ := file.Stat()
		//filePath := strings.TrimSuffix(file.Name(), info.Name())
		////fmt.Println("Filepath is : ", filePath)
		////fmt.Println("File name is: ", info.Name())
		//parts := strings.Split(info.Name(), ".")
		//extension := parts[len(parts)-1]
		//filename := strings.TrimSuffix(info.Name(), extension)
		////fmt.Println("Extension: ", extension)
		////fmt.Println("Filename: ", filename)
		//modFileName := slug.MakeLang(filename, "en") + "." + extension
		////fmt.Println("Modified filename is: ", modFileName)
		//newFilePath := filePath + modFileName
		//fmt.Println("New FilePath: ", newFilePath)
		//fmt.Println("Original FilePath: ", originalPath)
		//newRelativePath := strings.TrimPrefix(newFilePath, directory)
		//fmt.Println(newRelativePath)
		//fmt.Println(v)
		//v.Url = newRelativePath
		//v.FileName = modFileName
		//fmt.Println(v)
		//fmt.Println()
		//newDB.Save(&v)
		//os.Rename(originalPath, newFilePath)
	}

	//filepath.Walk( directory, func(path string, info os.FileInfo, err error) error {
	//	fmt.Println(slug.MakeLang(info.Name(), "en"))
	//	fmt.Println(path)
	//	fmt.Println(info.Name())
	//	fmt.Println(filepath.Dir(path))
	//	fmt.Println(filepath.Ext(path))
	//	fmt.Println(strings.TrimSuffix(info.Name(), filepath.Ext(path)))
	//	return nil
	//})
}
