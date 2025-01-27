package main
import (
    "fmt"
    "log"
    "sync"

    "github.com/rs/zerolog"
)

// CustomLogger represents a custom logger that wraps zerolog.Logger
type CustomLogger struct {
    logger *zerolog.Logger
}

// NewCustomLogger creates a new CustomLogger with a specified level
func NewCustomLogger(level zerolog.Level) *CustomLogger {
    zerolog.SetGlobalLevel(level)
    return &CustomLogger{
        logger: &zerolog.Logger{},
    }
}

// Info logs an information message with an extra field
func (cl *CustomLogger) Info(msg string, fieldName string, fieldValue interface{}) {
    cl.logger.Info().Str(fieldName, fmt.Sprintf("%v", fieldValue)).Msg(msg)
}

// Error logs an error message with an extra field
func (cl *CustomLogger) Error(msg string, fieldName string, fieldValue interface{}) {
    cl.logger.Error().Str(fieldName, fmt.Sprintf("%v", fieldValue)).Msg(msg)
}

func worker(id int, wg *sync.WaitGroup, cl *CustomLogger) {
    defer wg.Done()
    cl.Info("Worker started", "worker_id", id)
    // Simulate some work being done
    // Replace this with your actual work
    time.Sleep(time.Second)
    cl.Info("Worker finished", "worker_id", id)
}


func main() {
    numWorkers := 5
    var wg sync.WaitGroup