package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/koreset/homefnew/app/models"
	"github.com/koreset/homefnew/app/utils"
	"strings"
)

var newDB *gorm.DB
var homefDB *gorm.DB
var dbError error

func baseMigration() {

	var baseQuery = "select nid as id, title, type, created, changed as updated from node"
	rows, _ := homefDB.Raw(baseQuery).Rows()

	for rows.Next() {
		var post models.Post
		homefDB.ScanRows(rows, &post)
		newDB.Create(&post)
		newDB.Save(&post)
	}

	defer rows.Close()
}

func populateArticleBody() {
	var posts []models.Post
	newDB.Find(&posts)

	for _, v := range posts {
		rows, _ := homefDB.Raw("select body_value as body, body_summary as summary from field_data_body where entity_id = ? ", v.ID).Rows()

		for rows.Next() {
			homefDB.ScanRows(rows, &v)
		}

		v.Body = utils.CleanHtmlBody(v.Body)

		newDB.Save(&v)
	}
}

func populateImages() {
	var posts []models.Post
	newDB.Find(&posts)

	for _, v := range posts {

		rows, _ := homefDB.Raw("select entity_id as id, filename as file_name, uri as url from field_data_field_image, file_managed where field_image_fid = fid and entity_id = ?", v.ID).Rows()

		for rows.Next() {
			var image models.Image
			homefDB.ScanRows(rows, &image)
			image.Url = strings.Replace(image.Url, "public://", "", -1)
			v.Images = append(v.Images, image)
		}
		newDB.Save(&v)
	}

}

func populateVideoItems() {
	var posts []models.Post
	newDB.Find(&posts)
	for _, v := range posts {
		rows, _ := homefDB.Raw("select field_video_value as value, field_video_description_value as description from field_data_field_video as fdv, field_data_field_video_description as fvdd where fdv.entity_id = ? and fvdd.entity_id = ?", v.ID, v.ID).Rows()
		for rows.Next() {
			if v.Type == "video" {
				var video models.Video
				homefDB.ScanRows(rows, &video)
				fmt.Println(video)

				v.Body = video.Description
				v.Videos = append(v.Videos, video)
			}

			newDB.Save(&v)
		}
	}
}

func populateLinks() {
	var posts []models.Post
	newDB.Find(&posts)
	for _, v := range posts {
		rows, _ := homefDB.Raw("select field_url_url as url, field_url_title as title from field_data_field_url where entity_id = ?", v.ID).Rows()
		for rows.Next() {
			var link models.Link
			homefDB.ScanRows(rows, &link)
			fmt.Println(link)
			v.Links = append(v.Links, link)
		}

		newDB.Save(&v)
	}

}

func main() {
	//awsConnection := "homef:wordpass15@tcp(rds-mysql-homef.cb44dbuhyviz.eu-west-2.rds.amazonaws.com:3306)/homef?charset=utf8&parseTime=True&loc=Local"
	localConnection := "root:wordpass15@tcp(localhost:3306)/homef?charset=utf8&parseTime=True&loc=Local"
	newDB, dbError = gorm.Open("mysql", localConnection)
	if dbError != nil {
		panic(dbError)
	}

	homefDB, dbError = gorm.Open("mysql", "root:wordpass15@tcp(localhost:3306)/homef_db?charset=utf8&parseTime=True&loc=Local")
	if dbError != nil {
		panic(dbError)
	}

	newDB.LogMode(true)
	homefDB.LogMode(true)

	var post models.Post
	var video []models.Video
	var image []models.Image
	var link []models.Link

	newDB.Model(&post).Related(&video)
	newDB.Model(&post).Related(&image)
	newDB.Model(&post).Related(&link)
	newDB.AutoMigrate(&models.Post{}, &models.Video{}, &models.Image{}, &models.Link{})

	baseMigration()

	populateArticleBody()

	populateImages()

	populateVideoItems()

	populateLinks()

	//Populate article qu

	// for rows.Next() {
	// 	var result models.Content
	// 	homefDB.ScanRows(rows, &result)
	// 	newDB.Save(result)
	// 	// count++
	// 	// fmt.Println(count)
	// }
	// defer rows.Close()
	//var imageQuery string = `
	//select  field_data_field_image.bundle,
	//field_data_field_image.entity_id,
	//field_data_field_image.field_image_fid,
	//field_data_field_image.field_image_width,
	//field_data_field_image.field_image_height,
	//file_managed.filename,
	//file_managed.uri
	//from field_data_field_image, file_managed
	//where field_data_field_image.field_image_fid = file_managed.fid
	//`
	//rows, dbError := homefDB.Raw(imageQuery).Rows()
	//
	//if dbError != nil {
	//	panic(dbError)
	//}
	//
	//for rows.Next() {
	//	var media models.Media
	//	homefDB.ScanRows(rows, &media)
	//	media.Uri = strings.Replace(media.Uri, "public://", "", -1)
	//	newDB.Save(&media)
	//}

}
