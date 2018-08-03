package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/koreset/homef-gin/services"
)

func Home(c *gin.Context){
	payload := make(map[string] interface{})
	posts := services.GetPosts(0, 10)
	newsfeed := services.GetNewsFeed()
	videos := services.GetVideos()
	publications := services.GetPublications()
	payload["posts"] = posts
	payload["newsfeed"] = newsfeed
	payload["videos"] = videos
	payload["publications"] = publications

	c.HTML(http.StatusOK, "home_page", payload)

}