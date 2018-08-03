package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/koreset/homefnew/app/utils"
	"github.com/mmcdole/gofeed"
	"github.com/revel/modules/jobs/app/jobs"
	gr "github.com/revel/modules/orm/gorm/app"
	"github.com/revel/revel"
)

type FeedItem struct {
	gorm.Model
	Title string `gorm:"unique;not null"`
	Link  string `gorm:"not null;unique"`
}

func init() {
	fmt.Println("**** Starting FeedItem Crom job ****")
	revel.OnAppStart(func() {
		jobs.Every(60*time.Minute, FeedItem{})
	})
}

func (f FeedItem) Run() {
	feedUrl := "https://news.google.com/news/rss/search/section/q/Homef%20%22Nnimmo%20Bassey%22/Homef%20%22Nnimmo%20Bassey%22?hl=en&gl=ZA&ned=en_za"
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedUrl)
	if err != nil {
		panic(err)
	}

	for _, v := range feed.Items {
		var feedItem FeedItem
		fmt.Println(v.Title)
		fmt.Println(v.Link)
		feedItem = FeedItem{Title: v.Title, Link: v.Link}
		feedItem.CreatedAt, _ = utils.ParseDate(v.Published)
		feedItem.UpdatedAt, _ = utils.ParseDate(v.Published)

		gr.DB.Save(&feedItem)
	}
}
