package models

import "time"

type Cart struct {
	BuyerEmailId string
	ImageId int
}

type PreviousOrders struct {
	BuyerEmailId string
	SellerEmailId string
	Amount int
	BoughtAt  time.Time
}