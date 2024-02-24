package app

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/wmuga/wildberries-l0/config"
	"github.com/wmuga/wildberries-l0/internal/entity"
	"github.com/wmuga/wildberries-l0/internal/usecase"
)

func Flooder(cfg *config.Config) {
	nats, err := usecase.NewNatsOrderService(cfg.Nats)
	if err != nil {
		log.Fatalln(err)
		return
	}

	for {
		order := entity.GenerateOrder()
		if rand.Intn(3) == 0 {
			switch rand.Intn(3) {
			case 0:
				order.DeliveryID = 5
				fmt.Println("Changed delivery id")
			case 1:
				order.TrackNumber = "blahblah"
				fmt.Println("Changed track number")
			default:
				order.OrderUID = "blahblah"
				fmt.Println("Changed uid")
			}
		}
		nats.Publsish(order)
		time.Sleep(time.Second * 1)
	}
}
