package controller

import (
	"app/service"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Router is a collection of RESTful controllers
type Router interface {
	Start()
}

type router struct {
	router  *mux.Router
	service service.PaymentService
}

// NewControllers returns a fully configured router
func NewControllers(s service.PaymentService) Router {
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")

	controller := pamentsController{s}

	r.HandleFunc("/payment", controller.processPayment).Methods("POST")

	return router{
		service: s,
		router:  r,
	}
}

func (r router) Start() {
	srv := &http.Server{
		Handler:      r.router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
