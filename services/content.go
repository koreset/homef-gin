package services

import (
	"github.com/koreset/homefnew/app/models"

)

func GetContent(start, limit int) []models.Content {
	var results []models.Content
	GetDB().Preload("Media").Where("type in (?)", []string{"article", "press_release"}).Order("created desc").Offset(start).Limit(limit).Find(&results)
	return results
}

func GetLatestArticles(start, limit int) []models.Content {
	var results []models.Content
	GetDB().Where("type in (?)", []string{"article"}).Order("created desc").Limit(limit).Find(&results)
	return results
}
