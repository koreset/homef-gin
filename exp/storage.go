package main

import (
	"github.com/qor/oss/filesystem"
	"os"
	"fmt"
)

func main() {
	//storage := s3.New(s3.Config{AccessID: "access_id", AccessKey: "access_key", Region: "region", Bucket: "bucket", Endpoint: "cdn.getqor.com", ACL: awss3.BucketCannedACLPublicRead})
	wd, _ := os.Getwd()
	workingDir := wd + "/public/content"
	storage := filesystem.New(workingDir)


	// Save a reader interface into storage
	//storage.Put("/sample.txt", reader)

	// Get file with path
	//storage.Get("/sample.txt")

	// Get object as io.ReadCloser
	//storage.GetStream("/sample.txt")

	// Delete file with path
	//storage.Delete("/sample.txt")

	file, err := os.Open(workingDir + "/abnlogo.png")
	if err != nil {
		panic(err)
	}
	storage.Put("/sample.png", file)
	fmt.Println(workingDir)
	// List all objects under path
	objects, _ := storage.List("/field/image")

	fmt.Println(objects)

	for i, v := range objects {
		fmt.Println(i, v.Name)
	}

	// Get Public Accessible URL (useful if current file saved privately)
	//storage.GetURL("/sample.txt")
}
