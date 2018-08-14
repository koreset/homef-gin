package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koreset/homef-gin/services"
)

func Home(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPosts(0, 10)
	newsfeed := services.GetNewsFeed()
	videos := services.GetVideos()
	publications := services.GetPublications()
	payload["posts"] = posts
	payload["newsfeed"] = newsfeed
	payload["videos"] = videos
	payload["publications"] = publications
	payload["active"] = "home_page"

	c.HTML(http.StatusOK, "home_page", payload)

}

func Contacts(c *gin.Context) {
	posts := services.GetPosts(0, 10)

	c.HTML(http.StatusOK, "index2.html", posts)

}
