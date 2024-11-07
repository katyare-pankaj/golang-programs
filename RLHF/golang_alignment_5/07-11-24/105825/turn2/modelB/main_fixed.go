package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
)

// Task represents the work item that needs to be processed idempotently
type Task struct {
	ID      string `json:"id"`
	Payload string `json:"payload"`
}

// App is the main application struct
type App struct {
	rdb       *redis.Client
	mqConn    *amqp.Connection
	mqChannel *amqp.Channel
	mu        sync.Mutex
}

// NewApp creates a new instance of the App
func NewApp() *App {
	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Change this to your Redis server address
		Password: "",               // Set your Redis password if required
		DB:       0,                // Use default DB
	})

	// Initialize RabbitMQ connection and channel
	mqConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	mqChannel, err := mqConn.Channel()
	if err != nil {
		log.Fatalf("Failed to open RabbitMQ channel: %v", err)
	}

	// Declare the queue for idempotent tasks
	_, err = mqChannel.QueueDeclare(
		"idempotent_tasks", // name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
	}

	return &App{
		rdb:       rdb,
		mqConn:    mqConn,
		mqChannel: mqChannel,
	}
}

// handleRequest is the HTTP handler function that takes the request and enqueues it for idempotent processing
func (a *App) handleRequest(w http.ResponseWriter, r *http.Request) {
	// Decode the request payload into a Task struct
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if the idempotency key is already present in Redis
	if a.isProcessed(task.ID) {
		fmt.Fprintf(w, "Request with id '%s' has already been processed.\n", task.ID)
		return
	}

	// Enqueue the task for idempotent processing in RabbitMQ
	if err := a.enqueueTask(task); err != nil {
		http.Error(w, "Failed to enqueue task", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Request with id '%s' has been enqueued for processing.\n", task.ID)
}

// enqueueTask enqueues a task for idempotent processing in RabbitMQ
func (a *App) enqueueTask(task Task) error {
	// Marshal the task into JSON
	body, err := json.Marshal(task)
	if err != nil {
		return err
	}

	// Publish the task to the "idempotent_tasks" queue
	err = a.mqChannel.Publish(
		"",                 // exchange
		"idempotent_tasks", // routing key (queue name)
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish task to RabbitMQ: %w", err)
	}

	return nil
}
