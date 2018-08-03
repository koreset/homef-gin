package services

import (
	"fmt"

	"github.com/koreset/homefnew/app/models"
)

func GetNewsFeed() []models.FeedItem {
	var results []models.FeedItem

	GetDB().Limit(10).Find(&results)
	fmt.Println(results)
	return results
}
