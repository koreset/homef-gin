package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/koreset/homef-gin/services"
	"strconv"
)

func GetPost(c *gin.Context) {
	payload := make(map[string] interface{} )
	id := c.Param("id")
	fmt.Println("The ID: ", id)
	postID, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "content_not_found", nil)
		return
	}
	post := services.GetPost(postID)
	payload["post"] = post
	payload["active"] = "none"
	c.HTML(http.StatusOK, "single_post_new", payload)
}

func GetPublications(c *gin.Context){
	c.HTML(http.StatusOK, "publications_home", nil)
}