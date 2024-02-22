package main

import (
	"github.com/wmuga/wildberries-l0/internal/order"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=example password=example dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&order.Order{}, &order.Delivery{}, &order.Payment{}, &order.Item{})
	if err != nil {
		panic(err)
	}
}
