package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ratheesh-gopinadh-kumar/imageupload-go/db"
	"github.com/ratheesh-gopinadh-kumar/imageupload-go/handlers"
	"github.com/ratheesh-gopinadh-kumar/imageupload-go/models"
)

// func Init(){
// 	db.DBconnect()
// }

func main() {
	db.DBconnect()
	r := gin.Default()
	r.Static("/uploads","./uploads")//server uploaded image
	r.LoadHTMLGlob("template/*.html")//load html

	r.GET("/upload.html",func(c*gin.Context){
		c.HTML(200,"upload.html",gin.H{
			"title":"Home Page",
		})
	})


	//Define route for the fetching the image 
	r.POST("upload/image",handlers.UploadImage)
	
	// Define route for fetching image
	r.GET("/image/:id", handlers.ViewImage)
	r.Run(":3000")

}

// GetImageByID retrieves an image from the database by its ID
func GetImageByID(imageID string) (*models.Image, error) {
    var image models.Image
    if err := db.DB.Where("id = ?", imageID).First(&image).Error; err != nil {
        return nil, err
    }
    return &image, nil
}