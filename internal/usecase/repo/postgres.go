package repo

import (
	"fmt"

	"github.com/wmuga/wildberries-l0/config"
	"github.com/wmuga/wildberries-l0/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dsnFormat = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
)

// Postgres repo.
// Create instace with NewPostgresRepo
type postgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo(config config.DB) (OrderRepo, error) {
	// Open connection
	dsn := fmt.Sprintf(dsnFormat, config.Host, config.User, config.Password, config.Database, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	// Create tables
	err = db.AutoMigrate(&entity.Order{}, &entity.Delivery{}, &entity.Payment{}, &entity.Item{})
	if err != nil {
		return nil, err
	}
	return &postgresRepo{db}, nil
}

// Implementation of OrderRepo.AddOrder
func (p *postgresRepo) AddOrder(order *entity.Order) error {
	res := p.db.Create(order)
	return res.Error
}

// Implementation of OrderRepo.GetOrder
func (p *postgresRepo) GetOrder(id string) (*entity.Order, error) {
	order := &entity.Order{}

	// gorm.ErrRecordNotFound - не найдено
	res := p.db.Preload("Items").Preload("Delivery").Preload("Payment").First(order, "order_uid = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}

	return order, nil
}

// Implementation of OrderRepo.GetOrders
func (p *postgresRepo) GetOrders(count int) ([]entity.Order, error) {
	var orders []entity.Order
	res := p.db.Preload("Items").Preload("Delivery").Preload("Payment").Limit(count).Find(&orders)
	if res.Error != nil {
		return nil, res.Error
	}
	return orders, nil
}
