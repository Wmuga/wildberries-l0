package repo

import (
	"errors"

	"github.com/wmuga/wildberries-l0/internal/entity"
)

// Mock repo.
// Create instace with NewMockRepo
type mockRepo struct {
	orders     []entity.Order
	errorUID   string
	errorCount int
}

var (
	del1 = entity.Delivery{
		Name:    "Test",
		Phone:   "+9720000000",
		Zip:     "1231231",
		City:    "Test",
		Address: "test",
		Region:  "Test",
		Email:   "test@gmail.com",
	}
	order1 = entity.Order{
		OrderUID:    "123",
		TrackNumber: "Track1",
		Entry:       "WBIL",
		DeliveryID:  1,
		Delivery:    del1,
		Payment: entity.Payment{
			Transaction: "123",
			Currency:    "USD",
			Provider:    "wbpay",
			Amount:      213,
		},
		Items: []entity.Item{
			{
				TrackNumber: "Track1",
				Price:       123,
				TotalPrice:  321,
			},
		},
	}
	order2 = entity.Order{
		OrderUID:    "222",
		TrackNumber: "Track2",
		Entry:       "WBIL",
		DeliveryID:  1,
		Delivery:    del1,
		Payment: entity.Payment{
			Transaction: "222",
			Currency:    "USD",
			Provider:    "wbpay",
			Amount:      213,
		},
		Items: []entity.Item{
			{
				TrackNumber: "Track2",
				Price:       123,
				TotalPrice:  321,
			},
		},
	}
)

var (
	errMock = errors.New("mock error")
)

func NewMockRepo() OrderRepo {
	return &mockRepo{
		errorUID:   "-1",
		errorCount: -1,
		orders:     []entity.Order{order1, order2},
	}
}

// Implementation of OrderRepo.AddOrder
func (m *mockRepo) AddOrder(order *entity.Order) error {
	if order.OrderUID == m.errorUID {
		return errMock
	}
	m.orders = append(m.orders, *order)
	return nil
}

// Implementation of OrderRepo.GetOrder
func (m *mockRepo) GetOrder(id string) (*entity.Order, error) {
	if id == m.errorUID {
		return nil, errMock
	}

	m.orders[0].OrderUID = id
	m.orders[0].Payment.Transaction = id
	return &m.orders[0], nil
}

// Implementation of OrderRepo.GetOrders
func (m *mockRepo) GetOrders(count int) ([]entity.Order, error) {
	if count == m.errorCount {
		return nil, errMock
	}
	return m.orders[:count], nil
}
