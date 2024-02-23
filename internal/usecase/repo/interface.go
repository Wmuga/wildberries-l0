package repo

import "github.com/wmuga/wildberries-l0/internal/entity"

type OrderRepo interface {
	AddOrder(order *entity.Order) error
	GetOrder(id string) (*entity.Order, error)
	GetOrders(count int) ([]entity.Order, error)
}
