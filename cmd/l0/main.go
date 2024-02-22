package main

import (
	"log"

	"github.com/wmuga/wildberries-l0/config"
	"github.com/wmuga/wildberries-l0/internal/app"
	"github.com/wmuga/wildberries-l0/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	app.L0(cfg)

	dsn := "host=localhost user=wbl0user password=wbl0password dbname=wbl0 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&entity.Order{}, &entity.Delivery{}, &entity.Payment{}, &entity.Item{})
	if err != nil {
		panic(err)
	}
}
