package usecase

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/wmuga/wildberries-l0/config"
	"github.com/wmuga/wildberries-l0/internal/entity"
)

// Service connection to nats-streaming.
// Create new instance with NewNatsService
type NatsOrderService struct {
	connection *nats.Conn
	subject    string
}

// Creates new connection to nats-streaming
func NewNatsOrderService(cfg config.Nats) (*NatsOrderService, error) {
	nc, err := nats.Connect(cfg.URL)
	if err != nil {
		return nil, err
	}

	return &NatsOrderService{nc, cfg.Subject}, nil
}

// Subscribes to order's topic. Calls callback on every message
func (n *NatsOrderService) Subscibe(callback func(*entity.Order, error)) {
	n.connection.Subscribe(n.subject, func(msg *nats.Msg) {
		order := &entity.Order{}
		err := json.Unmarshal(msg.Data, order)
		if err != nil {
			callback(nil, err)
			return
		}
		callback(order, order.Verify())
	})
}

// Published order to topic
func (n *NatsOrderService) Publsish(order *entity.Order) error {
	message, err := json.Marshal(order)
	if err != nil {
		return err
	}
	n.connection.Publish(n.subject, message)
	return nil
}
