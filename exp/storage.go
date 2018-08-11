package main

import (
	"os"
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
	"github.com/qor/media"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"net/url"
	"strings"
	"path/filepath"
	"net/http"
	"io"
)

type TestPost struct {
	gorm.Model
	Images []ContentFile
}

type ContentFile struct {
	gorm.Model
	Document oss.OSS `sql:"size:4294967295;" media_library:"url:/publications/{{basename}}.{{extension}};path:./public"`
	//Document oss.OSS
	TestPostID uint
}

//func getFuncMap(scope *gorm.Scope, field *gorm.Field, filename string) template.FuncMap {
//	hash := func() string { return strings.Replace(time.Now().Format("20060102150506.000000000"), ".", "", -1) }
//	return template.FuncMap{
//		"class":       func() string { return inflection.Plural(utils.ToParamString(scope.GetModelStruct().ModelType.Name())) },
//		"primary_key": func() string { return fmt.Sprintf("%v", scope.PrimaryKeyValue()) },
//		"column":      func() string { return strings.ToLower(field.Name) },
//		"filename":    func() string { return filename },
//		"basename":    func() string { return strings.TrimSuffix(path.Base(filename), path.Ext(filename)) },
//		"hash":        hash,
//		"filename_with_hash": func() string {
//			return urlReplacer.ReplaceAllString(fmt.Sprintf("%s.%v%v", slug.Make(strings.TrimSuffix(path.Base(filename), path.Ext(filename))), hash(), path.Ext(filename)), "-")
//		},
//		"extension": func() string { return strings.TrimPrefix(path.Ext(filename), ".") },
//	}


func main() {
	localConnection := "root:wordpass15@tcp(localhost:3306)/stores?charset=utf8&parseTime=True&loc=Local"
	newDB, dbError := gorm.Open("mysql", localConnection)
	if dbError != nil {
		panic(dbError)
	}


	newDB.AutoMigrate(&TestPost{}, &ContentFile{})
	media.RegisterCallbacks(newDB)


	workingDir := "/Users/jome/projects/homef/files/pubs/"
	myfile, _ := os.Open(workingDir + "eco-instigatorv2.pdf")
	var newpost TestPost
	var images []ContentFile

	newDB.Model(&newpost).Related(&images)
	image := ContentFile{}
	image.Document.Scan(myfile)
	fmt.Println(image.Document.Url)
	images = append(images, image)
	newpost = TestPost{
		Images: images,
	}

	newDB.Save(&newpost)

	//storage := s3.New(s3.Config{AccessID: "access_id", AccessKey: "access_key", Region: "region", Bucket: "bucket", Endpoint: "cdn.getqor.com", ACL: awss3.BucketCannedACLPublicRead})
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


func openFileByURL(rawURL string) (*os.File, error) {
	if fileURL, err := url.Parse(rawURL); err != nil {
		return nil, err
	} else {
		path := fileURL.Path
		segments := strings.Split(path, "/")
		fileName := segments[len(segments)-1]

		filePath := filepath.Join(os.TempDir(), fileName)

		if _, err := os.Stat(filePath); err == nil {
			return os.Open(filePath)
		}

		file, err := os.Create(filePath)
		if err != nil {
			return file, err
		}

		check := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}
		resp, err := check.Get(rawURL) // add a filter to check redirect
		if err != nil {
			return file, err
		}
		defer resp.Body.Close()
		fmt.Printf("----> Downloaded %v\n", rawURL)

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return file, err
		}
		return file, nil
	}
}