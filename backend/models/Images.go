package models

import (
	"gorm.io/gorm"
	"time"
)



type Images struct {
	gorm.Model
	EmailId string
	ImageId int
	Title string
	Description string
	Price float32
	UploadedAt time.Time
	ImageURL string
	WImageURL string
}