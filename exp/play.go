package main

import (
	"fmt"
	"github.com/koreset/homef-gin/services"
)

func main() {
	videos := services.GetVideos()
	fmt.Println(videos)
}
