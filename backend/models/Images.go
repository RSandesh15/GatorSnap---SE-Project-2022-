package models

import (
	"time"
)

type Image struct {
	EmailId     string
	ImageId     int `gorm:"primaryKey"`
	Title       string
	Description string
	Price       float32
	UploadedAt  time.Time
	ImageURL    string
	WImageURL   string
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
