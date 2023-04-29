package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/virsel/SP-verteilte-Systeme/event"
)

const (
	address = "localhost:8080"
)

// createEvent calls the RPC method CreateEvent of EventServer
func createEvent(client pb.EventClient, Event *pb.EventRequest) {
	resp, err := client.CreateEvent(context.Background(), Event)
	if err != nil {
		log.Fatalf("Could not create Event: %v", err)
	}
	if resp.Success {
		log.Printf("Event added id: %s", resp.Id)
	}
}

// getEvents calls the RPC method GetEvents of EventServer
func getEvents(client pb.EventClient, filter *pb.GetEventFilter) {
	// calling the streaming API
	resp, err := client.GetEvent(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get Event: %v", err)
	}
	log.Printf("Get Event id %s: %s", filter.Id, resp.Name)
}

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := pb.NewEventClient(conn)

	event := &pb.EventRequest{
		Name: "order1",
	}

	// Create a new event
	createEvent(client, event)

	event = &pb.EventRequest{
		Name: "order2",
	}

	// Create a new event
	createEvent(client, event)

	// Filter with an empty Keyword
	filter := &pb.GetEventFilter{Id: "cc0c1003-4461-4b96-b5d6-5e111ba441f7"}
	getEvents(client, filter)
}
