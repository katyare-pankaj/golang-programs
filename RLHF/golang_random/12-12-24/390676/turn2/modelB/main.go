package main

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron"
)

// Task represents a periodic task
type Task struct {
	name string
	spec string
	fn   func()
}

// TaskScheduler manages the execution of periodic tasks
type TaskScheduler struct {
	ctx     context.Context
	cancel  context.CancelFunc
	c       *cron.Cron
	tasks   []Task
	errorCh chan error
}

// NewTaskScheduler creates a new task scheduler
func NewTaskScheduler() *TaskScheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &TaskScheduler{
		ctx:     ctx,
		cancel:  cancel,
		c:       cron.New(),
		errorCh: make(chan error),
	}
}

// AddTask adds a new periodic task
func (ts *TaskScheduler) AddTask(name, spec string, fn func()) {
	task := Task{name: name, spec: spec, fn: fn}
	ts.tasks = append(ts.tasks, task)
	ts.c.AddFunc(spec, func() {
		select {
		case <-ts.ctx.Done():
			return
		default:
		}
		fn()
	})
}

// Start starts the scheduler
func (ts *TaskScheduler) Start() {
	go ts.c.Start()
	go ts.listenForErrors()
}

// Stop stops the scheduler
func (ts *TaskScheduler) Stop() {
	ts.cancel()
	ts.c.Stop()
	close(ts.errorCh)
}

// listenForErrors listens for errors and logs them
func (ts *TaskScheduler) listenForErrors() {
	for err := range ts.errorCh {
		log.Printf("Error: %v", err)
	}
}

func main() {
	scheduler := NewTaskScheduler()

	scheduler.AddTask("Every Minute Task", "@every 1m", func() {
		log.Println("Every Minute Task executed")
	})

	scheduler.AddTask("Every Hour Task", "@every 1h", func() {
		log.Println("Every Hour Task executed")
	})

	scheduler.Start()

	defer scheduler.Stop()

	select {
	case <-time.After(10 * time.Second):
		log.Println("Scheduler running for 10 minutes, stopping now")
	}
}
