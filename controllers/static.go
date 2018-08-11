package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AboutUs(c *gin.Context){
	c.HTML(http.StatusOK, "about_us", nil)
}
