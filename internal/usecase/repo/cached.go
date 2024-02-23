package repo

import (
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/wmuga/wildberries-l0/config"
	"github.com/wmuga/wildberries-l0/internal/entity"
)

var (
	ErrCantConvert = errors.New("can't convert interface{} to *entity.Order")
)

// Caches OrderRepo requests.
// Create instace with NewCachedRepo
type cachedRepo struct {
	repo  OrderRepo
	cache *cache.Cache
}

// Creates new instance of cached repo
func NewCachedRepo(repo OrderRepo, config config.Cache) (OrderRepo, error) {
	c := &cachedRepo{
		repo,
		cache.New(time.Minute*time.Duration(config.TTLMins), time.Minute*time.Duration(config.PutgeMins)),
	}

	// Restore cache from database
	orders, err := repo.GetOrders(config.RestoreSize)
	if err != nil {
		return nil, err
	}
	for _, order := range orders {
		err = c.cache.Add(order.OrderUID, &order, cache.DefaultExpiration)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// Implementation of OrderRepo.AddOrder
func (c *cachedRepo) AddOrder(order *entity.Order) error {
	// Add to repo
	err := c.repo.AddOrder(order)
	if err != nil {
		return err
	}
	// Add to cache
	err = c.cache.Add(order.OrderUID, order, cache.DefaultExpiration)
	if err != nil {
		return err
	}
	return nil
}

// Implementation of OrderRepo.GetOrder
func (c *cachedRepo) GetOrder(id string) (*entity.Order, error) {
	// Check in cache
	orderInterface, found := c.cache.Get(id)
	if !found {
		// Retrive from repo
		order, err := c.repo.GetOrder(id)
		if err != nil {
			return nil, err
		}
		// Add to cache
		err = c.cache.Add(order.OrderUID, order, cache.DefaultExpiration)
		if err != nil {
			return nil, err
		}
		return order, nil
	}
	// Convert to pointer
	order, conv := orderInterface.(*entity.Order)
	if !conv {
		return nil, ErrCantConvert
	}
	return order, nil
}

// Implementation of OrderRepo.GetOrders
func (c *cachedRepo) GetOrders(count int) ([]entity.Order, error) {
	// No functional "Get many" in cache. Just passthrough
	return c.repo.GetOrders(count)
}
