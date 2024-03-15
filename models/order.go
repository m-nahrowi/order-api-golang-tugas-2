package models

import "time"

type Order struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID"`
	Status       string    `json:"status"`
}

type Item struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	OrderID     uint   `json:"-"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
