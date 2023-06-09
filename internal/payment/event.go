package payment

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	pb "github.com/virsel/SP-verteilte-Systeme/event"
	"github.com/virsel/SP-verteilte-Systeme/internal/model"
)

const (
	streamSubjectsname = "ORDERS.new"
)

type Server struct {
	pb.UnimplementedEventServer
	Nats nats.JetStreamContext
}

func (s *Server) ConsumeEvent(js nats.JetStreamContext) {
	_, err := js.Subscribe(streamSubjectsname, func(m *nats.Msg) {
		err := m.Ack()

		if err != nil {
			log.Println("Unable to Ack", err)
			return
		}

		var orderEvt model.OrderEvent

		err = json.Unmarshal(m.Data, &orderEvt)
		if err != nil {
			log.Panic(err)
		}

		log.Printf("Consumer  =>  Subject: %s  -  ID: %s  -  Name: %s\n", m.Subject, orderEvt.Id, orderEvt.Name)
	})

	if err != nil {
		log.Println("Subscribe failed")
		return
	}
}
