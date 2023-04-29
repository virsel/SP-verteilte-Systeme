package main

import (
	"log"
	"net"

	pb "github.com/virsel/SP-verteilte-Systeme/event"
	"github.com/virsel/SP-verteilte-Systeme/internal/payment"
	"github.com/virsel/SP-verteilte-Systeme/pkg/stream"
	"github.com/virsel/SP-verteilte-Systeme/pkg/utils"
	"google.golang.org/grpc"
)

const (
	port       = "50051"
	streamName = "ORDERS"
)

func main() {

	js, err := stream.JetStreamInit(streamName)
	if err != nil {
		log.Println(err)
		return
	}

	port := utils.GetEnv("PORT", port)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()

	server := &payment.Server{
		Nats: js,
	}
	pb.RegisterEventServer(s, server)

	server.ConsumeEvent(js)

	log.Printf("gRPC Server listening on %v", ":"+port)
	s.Serve(lis)
}
