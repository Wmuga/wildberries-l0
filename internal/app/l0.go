package app

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/wmuga/wildberries-l0/config"
	"github.com/wmuga/wildberries-l0/internal/controllers"
	"github.com/wmuga/wildberries-l0/internal/entity"
	"github.com/wmuga/wildberries-l0/internal/middleware"
	"github.com/wmuga/wildberries-l0/internal/usecase"
)

func L0(config *config.Config) {
	nats, err := usecase.NewNatsOrderService(config.Nats)
	if err != nil {
		log.Fatalln(err)
	}

	orders, err := usecase.NewOrderService(config)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Order service ready")
	// adds orders from nats to databse
	err = nats.Subscibe(func(o *entity.Order, err error) {
		if err != nil {
			log.Println(err)
			return
		}
		if err = orders.AddOrder(o); err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Nats service ready")
	// add order from json if not exists
	addFromJSON(orders)
	// router for ui
	r := mux.NewRouter()
	controllers.NewOrdersRouter(r, orders, log.New(os.Stdout, "[UI]", log.LUTC))
	r.Use(middleware.GetRequestLogger(log.New(os.Stdout, "[REQ]", log.LUTC)))
	log.Println("Listening at", config.HTTP.URL)
	log.Println(http.ListenAndServe(config.HTTP.URL, r))
}

func addFromJSON(orders *usecase.OrderService) {
	if _, err := orders.GetOrder("b563feb7b2b84b6test"); err == nil {
		return
	}

	bytes, err := os.ReadFile("model.json")
	if err != nil {
		log.Fatalln(err)
	}

	order := &entity.Order{}
	err = json.Unmarshal(bytes, order)
	if err != nil {
		log.Fatalln(err)
	}

	err = orders.AddOrder(order)
	if err != nil {
		log.Fatalln(err)
	}
}
