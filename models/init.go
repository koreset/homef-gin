package models

import (
	"net/http"

	"github.com/qor/admin"
	gorm "github.com/revel/modules/orm/gorm/app"
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(InitDB)
	revel.OnAppStart(installHandlers)
}

func InitDB() {
	gorm.InitDB()
	gorm.DB.AutoMigrate(&Post{}, &Video{}, &Image{}, &Link{}, &FeedItem{})

}

func installHandlers() {
	revel.AddInitEventHandler(func(event int, _ interface{}) (r int) {
		if event == revel.ENGINE_STARTED {
			var (
				serveMux     = http.NewServeMux()
				revelHandler = revel.CurrentEngine.(*revel.GoHttpServer).Server.Handler
			)
			// Initalize
			Admin := admin.New(&admin.AdminConfig{DB: gorm.DB})
			Admin.AddResource(&Content{})
			Admin.MountTo("/admin", serveMux)
			serveMux.Handle("/", revelHandler)
			revel.CurrentEngine.(*revel.GoHttpServer).Server.Handler = serveMux

		}
		return
	})
}
