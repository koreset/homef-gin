package main

import (
	"fmt"
	"github.com/koreset/homefnew/app/services"
)

func main() {
	videos := services.GetVideos()
	fmt.Println(videos)
}
