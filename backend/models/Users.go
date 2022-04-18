package models

type Buyer struct {
	buyerId      int    `gorm:"primaryKey" json:"buyerId"`
	buyerEmailId string `json:"buyerEmailId"`
	Name         string `json:"Name"`
}

type Seller struct {
	sellerID      int
	sellerEmailId string
	Name          string
}
