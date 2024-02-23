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
	"github.com/wmuga/wildberries-l0/internal/usecase"
)

func L0(config *config.Config) {
	orders, err := usecase.NewOrderService(config)
	if err != nil {
		log.Fatalln(err)
	}
	// add order from json if not exists
	addFromJson(orders)
	r := mux.NewRouter()
	controllers.NewOrdersRouter(r, orders, log.New(os.Stdout, "[API]", log.LUTC))

	log.Println("Listening at", config.HTTP.URL)
	log.Println(http.ListenAndServe(config.HTTP.URL, r))
}

func addFromJson(orders *usecase.OrderService) {
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
