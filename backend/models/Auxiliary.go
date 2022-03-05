package models

import "time"

type Cart struct {
	BuyerEmailId string
	ImageId int
}

type PreviousOrders struct {
	BuyerEmailId string
	ImageId int
	BoughtAt  time.Time
}