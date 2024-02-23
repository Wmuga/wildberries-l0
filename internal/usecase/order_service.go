package usecase

import (
	"github.com/wmuga/wildberries-l0/config"
	"github.com/wmuga/wildberries-l0/internal/entity"
	"github.com/wmuga/wildberries-l0/internal/usecase/repo"
)

type OrderService struct {
	db     repo.OrderRepo
	cached repo.OrderRepo
}

func NewOrderService(cfg *config.Config) (*OrderService, error) {
	db, err := repo.NewPostgresRepo(cfg.DB)
	if err != nil {
		return nil, err
	}

	cache, err := repo.NewCachedRepo(db, cfg.Cache)
	if err != nil {
		return nil, err
	}

	return &OrderService{
		db,
		cache,
	}, nil
}

func (s *OrderService) AddOrder(order *entity.Order) error {
	return s.db.AddOrder(order)
}

func (s *OrderService) GetOrder(id string) (*entity.Order, error) {
	return s.cached.GetOrder(id)
}
