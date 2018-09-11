package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/homef-gin/services"
	"net/http"
)

func GetTest(c *gin.Context){
	payload := make(map[string] interface{} )

	posts := services.GetPosts(0, 4)
	payload["posts"] = posts


	c.HTML(http.StatusOK, "testing", payload)
}
