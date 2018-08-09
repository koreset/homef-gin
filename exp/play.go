package main

import (
	"fmt"
	"github.com/koreset/homef-gin/services"
)

type User struct {
	Name string
}
func main() {
	videos := services.GetVideos()
	fmt.Println(videos)


}
