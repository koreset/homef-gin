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
	"flag"
	"fmt"
	"github.com/qor/admin"
	"net/http"
)

var db *gorm.DB
var funcMaps template.FuncMap


func SetupRouter() *gin.Engine {
	mux := http.NewServeMux()

	Admin := admin.New(&admin.AdminConfig{DB: db})

	Admin.MountTo("/admin", mux)

	Admin.AddResource(&models.Content{})

	router := gin.Default()
	router.SetFuncMap(setupTemplatFuncs())
	router.LoadHTMLGlob("views/**/*")


	router.GET("/", controllers.Home)


	router.Static("/public", "./public")
	router.Any("/admin/*resources", gin.WrapH(mux))

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
	port := flag.String("port", "4000", "The port the app will listen to")
	host := flag.String("host", "0.0.0.0", "The ip address to listen on")
	flag.Parse()

	SetupDB()
	defer db.Close()
	r := SetupRouter()
	fmt.Println(*host, *port)
	r.Run(fmt.Sprintf("%s:%s", *host, *port))
}
