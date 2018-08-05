package services

import (
	"fmt"

	"github.com/koreset/homef-gin/models"
)

func GetNewsFeed() []models.FeedItem {
	var results []models.FeedItem

	GetDB().Limit(10).Find(&results)
	fmt.Println(results)
	return results
}
