package order

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID                int       `json:"-" gorm:"primaryKey,autoIncrement"`
	OrderUID          string    `json:"order_uid" gorm:"unique"`
	TrackNumber       string    `json:"track_number" gorm:"unique"`
	Entry             string    `json:"entry"`
	DeliveryID        int       `json:"-" gorm:"unique"`
	Delivery          Delivery  `json:"delivery" gorm:"foreignKey:DeliveryID;references:DeliveryID"`
	Payment           Payment   `json:"payment" gorm:"foreignKey:Transaction;references:OrderUID"`
	Items             []Item    `json:"items" gorm:"foreignKey:TrackNumber;references:TrackNumber"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type Delivery struct {
	gorm.Model
	DeliveryID int    `json:"-" gorm:"autoIncrement"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Zip        string `json:"zip"`
	City       string `json:"city"`
	Address    string `json:"address"`
	Region     string `json:"region"`
	Email      string `json:"email"`
}

type Payment struct {
	gorm.Model
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDT    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	gorm.Model
	ChrtID      int    `json:"chrt_id" gorm:"primaryKey"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}
