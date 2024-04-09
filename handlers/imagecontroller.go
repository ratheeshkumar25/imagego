package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path"

	//"strconv"
	//"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ratheesh-gopinadh-kumar/imageupload-go/db"
	"github.com/ratheesh-gopinadh-kumar/imageupload-go/models"
	//"github.com/ratheesh-gopinadh-kumar/imageupload-go/models"
)

const (
	maxFileSize = 2 * 1024 * 1024
	imageDir    = "uploads/images"
)

// UploadImage handles the image upload functionality
func UploadImage(c *gin.Context) {
    file, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate file size
    if file.Size > maxFileSize {
        c.JSON(http.StatusBadRequest, gin.H{"error": "max image size is 2mb"})
        return
    }

    // Generate unique file name
    now := time.Now().UTC()
    unixTime := now.UnixNano()
    fileExtension := path.Ext(file.Filename)
    newFileName := fmt.Sprintf("%d%s", unixTime, fileExtension)

    // Create directory if not exists
    storingPath := path.Join(imageDir, now.Format("2006/01/02"))
    if err := os.MkdirAll(storingPath, os.ModePerm); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
        return
    }

    // Save the uploaded file
    filePath := path.Join(storingPath, newFileName)
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    // Create and save image model in database
    image := models.Image{
        FileName:    filePath, // Save full file path
        ContentType: file.Header.Get("Content-Type"),
        Size:        file.Size,
        UploadDate:  now,
    }
    if err := db.DB.Create(&image).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image data"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "image": image})
}

// Handle the view image functionality
func ViewImage(c *gin.Context) {
    imageID := c.Param("id")

    // Retrieve image data from the database
    imagePath, err := db.GetImageByID(imageID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
        return
    }

    // Render HTML page to display the image
    c.HTML(http.StatusOK, "view_image.html", gin.H{
        "title": "View Image",
        "image": imagePath,
    })
}

