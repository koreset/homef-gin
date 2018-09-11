package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/koreset/gtf"
	"github.com/koreset/homef-gin/controllers"
	"github.com/koreset/homef-gin/models"
	"github.com/koreset/homef-gin/services"
	"github.com/koreset/homef-gin/utils"
	"github.com/qor/admin"
	"github.com/qor/media"
	"github.com/qor/media/asset_manager"
	"github.com/qor/media/media_library"
)

var db *gorm.DB
var funcMaps template.FuncMap

// AutoMigrate run auto migration
func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.AutoMigrate(value)
	}
}

func SetupRouter() *gin.Engine {
	mux := http.NewServeMux()

	Admin := admin.New(&admin.AdminConfig{DB: db})

	Admin.MountTo("/admin", mux)

	assetManager := Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})
	// Add Media Library
	Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Menu: []string{"Site Management"}})

	post := Admin.AddResource(&models.Post{}, &admin.Config{Name: "Posts", Menu: []string{"Content Management"}})
	post.IndexAttrs("ID", "Title", "Body", "Summary", "Images", "Videos", "Links", "Type")
	post.NewAttrs("Title", "Body", "Summary", "Images", "Videos", "Links", "Type")
	post.Meta(&admin.Meta{Name: "Body", Config: &admin.RichEditorConfig{AssetManager: assetManager}})

	router := gin.Default()
	router.SetFuncMap(setupTemplatFuncs())
	router.LoadHTMLGlob("views/**/*")

	router.GET("/", controllers.Home)
	router.GET("/aboutus", controllers.AboutUs)
	router.GET("/contacts", controllers.Contacts)
	router.GET("posts/:id", controllers.GetPost)
	router.GET("publications", controllers.GetPublications)
	router.GET("/test", controllers.GetTest)

	router.Static("/public", "./public")
	router.Any("/admin/*resources", gin.WrapH(mux))
	router.NoRoute(func(context *gin.Context) {
		fmt.Println(">>>>>>>>>>>>>>>>>> 404 <<<<<<<<<<<<<<<<<<<")
		context.HTML(http.StatusNotFound, "content_not_found", nil)
	})
	return router
}

func setupTemplatFuncs() template.FuncMap {
	funcMaps = make(template.FuncMap)
	funcMaps["unsafeHtml"] = utils.UnsafeHtml
	funcMaps["stripSummaryTags"] = utils.StripSummaryTags
	funcMaps["displayDateString"] = utils.DisplayDateString
	funcMaps["displayDate"] = utils.DisplayDateV2
	funcMaps["truncateBody"] = utils.TruncateBody

	gtf.Inject(funcMaps)
	return funcMaps
}

func SetupDB() {
	db = services.Init()
	db.AutoMigrate(&models.Post{}, &models.Video{}, &models.Image{}, &models.Link{}, &models.FeedItem{})
	media.RegisterCallbacks(db)
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
