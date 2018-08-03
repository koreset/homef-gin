package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/koreset/homef-gin/services"
	"github.com/jinzhu/gorm"
	"github.com/koreset/homef-gin/models"
	"github.com/koreset/homef-gin/utils"
	"html/template"
	"github.com/koreset/gtf"
	"github.com/koreset/homef-gin/controllers"
)

var db *gorm.DB
var funcMaps template.FuncMap


func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.SetFuncMap(setupTemplatFuncs())
	router.LoadHTMLGlob("views/**/*")


	router.GET("/", controllers.Home)


	router.Static("/public", "./public")

	return router
}

func setupTemplatFuncs() template.FuncMap{
	funcMaps = make(template.FuncMap)
	funcMaps["unsafeHtml"] = utils.UnsafeHtml
	funcMaps["stripSummaryTags"] = utils.StripSummaryTags
	funcMaps["displayDateString"] = utils.DisplayDateString
	funcMaps["displayDate"] = utils.DisplayDateV2

	gtf.Inject(funcMaps)
	return funcMaps
}

func SetupDB() {
	db = services.Init()
	db.AutoMigrate(&models.Content{})
}

func main() {
	SetupDB()
	defer db.Close()
	r := SetupRouter()
	r.Run(":3000")
}
