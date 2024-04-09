package models

import "time"

// Image represents the metadata of an uploaded image
type Image struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	FileName    string    `json:"file_name"`
	Size        int64     `json:"size"`
	ContentType string    `json:"content_type"`
	UploadDate  time.Time `json:"upload_date"`
}

// TableName specifies the database table name for the Image model
func (Image) TableName() string {
	return "images"
	
}