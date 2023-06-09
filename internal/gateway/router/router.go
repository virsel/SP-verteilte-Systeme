package router

import (
	"github.com/gorilla/mux"
	"github.com/virsel/SP-verteilte-Systeme/internal/gateway/handler"
	"github.com/virsel/SP-verteilte-Systeme/internal/gateway/service"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

func InitRouter(router *mux.Router) {

	repos := &service.SrvRepository{}
	mws := &service.Middleware{}

	h := handler.NewHandler(repos)

	r := router.PathPrefix("/api/gw").Subrouter()
	r.Use(otelmux.Middleware("gateway"))

	r.HandleFunc("/", h.HealthEndpoint).Methods("GET")
	r.HandleFunc("/order/create", mws.Middleware(h.CreateOrder)).Methods("POST")
	r.HandleFunc("/order/get/{id}", mws.Middleware(h.GetOrder)).Methods("GET")
}
