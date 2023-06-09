package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	pb "github.com/virsel/SP-verteilte-Systeme/event"
	"github.com/virsel/SP-verteilte-Systeme/internal/gateway/service"
	"github.com/virsel/SP-verteilte-Systeme/internal/model"
	"github.com/virsel/SP-verteilte-Systeme/pkg/grpcutil"
	"github.com/virsel/SP-verteilte-Systeme/pkg/utils"
)

const (
	address = "localhost:8080"
)

type Handler struct {
	repository service.Repository
}

func NewHandler(sr service.Repository) *Handler {
	return &Handler{
		repository: sr,
	}
}

// HealthEndpoint for srv status
func (h *Handler) HealthEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service Healthy"))
}

// Create Order
func (h *Handler) CreateOrder(w http.ResponseWriter, req *http.Request) {

	var order model.OrderEvent
	_ = json.NewDecoder(req.Body).Decode(&order)

	if order.Name == "" {
		err := "error missing name"
		log.Println(err)
		http.Error(w, err, http.StatusBadRequest)
		return
	}

	conn := grpcutil.GetgrpcConn(utils.GetEnv("ORDER_GRPC_ADDR", address))
	defer conn.Close()

	client := pb.NewEventClient(conn)

	event := &pb.EventRequest{
		Name: order.Name,
	}
	err := h.repository.CreateOrder(client, event)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// Get Order
func (h *Handler) GetOrder(w http.ResponseWriter, req *http.Request) {

	reqId := mux.Vars(req)["id"]
	if reqId == "" {
		log.Println("Id cannot be empty")
	}

	conn := grpcutil.GetgrpcConn(utils.GetEnv("ORDER_GRPC_ADDR", address))
	defer conn.Close()

	client := pb.NewEventClient(conn)

	// Filter with an empty Keyword
	filter := &pb.GetEventFilter{Id: reqId}
	resOrder, err := h.repository.GetOrder(client, filter)
	if err != nil {
		log.Println("GetOrder", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resOrder)
	}
}
