package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"interview-rest/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	m := &event.CommandMonitor{
		Started: func(_ context.Context, e *event.CommandStartedEvent) {
			log.Println(fmt.Sprintf("Command: %v", e.Command))
		},
		Succeeded: func(_ context.Context, e *event.CommandSucceededEvent) {
			log.Println(fmt.Sprintf("Succeeded: %v", e.Reply))
		},
		Failed: func(_ context.Context, e *event.CommandFailedEvent) {
			log.Println(fmt.Sprintf("Succeeded: %v", e.Failure))
		},
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetMonitor(m) // replace with your MongoDB URI
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	defer client.Disconnect(context.TODO())
	server := server.NewServer(8080, client)

	server.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	sig := <-sigChan

	fmt.Printf("Received signal: %s. Exiting...\n", sig)
}
