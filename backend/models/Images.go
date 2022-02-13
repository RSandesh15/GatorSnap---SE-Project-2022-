package models

import (
	"time"
)

type Image struct {
	EmailId     string
	ImageId     int       `gorm:"primaryKey" json:"imageId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	UploadedAt  time.Time `json:"uploadedAt"`
	ImageURL    string    
	WImageURL   string    `json:"wImageUrl"`
}

type ProductCatalogue struct {
	ImageId   int      `json:"imageId"`
	Price     float32  `json:"price"`
	Title     string   `json:"title"`
	WImageURL string   `json:"wImageUrl"`
	Genre     []string `json:"genres"`
}

type Genre struct {
	ImageId   int
	GenreType string
}

type GenreCategories struct {
	Category string `gorm:"primaryKey"`
}
