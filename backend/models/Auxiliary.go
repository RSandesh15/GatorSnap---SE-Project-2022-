package models

import "time"

type Cart struct {
	BuyerEmailId string
	ImageId int
}

type PreviousOrders struct {
	BuyerEmailId string
	SellerEmailId string
	ImageId int
	Title string
	Price float32
	BoughtAt  time.Time
}