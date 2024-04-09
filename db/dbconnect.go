package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ratheesh-gopinadh-kumar/imageupload-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() {
    //Load env file 
	err := godotenv.Load()
	if err != nil {
		panic("failed to connect to env")
	}
	//connect to dd with credentials provide in db 
	dsn := os.Getenv("DB")
	fmt.Println("DSN:", os.Getenv("DB"))

	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil{
		panic("Failed to connect database")
	}
	DB = db
    // Automigrate the entities 
	DB.AutoMigrate(&models.Image{})


}

// GetImageByID retrieves an image from the database by its ID
func GetImageByID(imageID string) (*models.Image, error) {
    var image models.Image
    if err := DB.Where("id = ?", imageID).First(&image).Error; err != nil {
        return nil, err
    }
    return &image, nil
}