package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wmuga/wildberries-l0/internal/usecase"
	"gorm.io/gorm"
)

type ordersRouter struct {
	orders    *usecase.OrderService
	logger    *log.Logger
	templates *template.Template
}

func NewOrdersRouter(r *mux.Router, orders *usecase.OrderService, logger *log.Logger) {
	templates := template.Must(template.ParseGlob("web/templates/*.html"))

	api := &ordersRouter{
		orders,
		logger,
		templates,
	}

	r.HandleFunc("/order/{id}", api.GetOrder).Methods("GET")
}

func (api *ordersRouter) GetOrder(w http.ResponseWriter, r *http.Request) {
	id, ex := mux.Vars(r)["id"]
	if !ex {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	order, err := api.orders.GetOrder(id)
	if err != nil && err == gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		api.logger.Println(err)
		return
	}

	err = api.templates.ExecuteTemplate(w, "order.html", order)
	if err != nil {
		api.logger.Println(err)
	}
}
