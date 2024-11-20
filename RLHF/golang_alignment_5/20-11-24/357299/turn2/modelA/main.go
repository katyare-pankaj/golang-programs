// main.go
package main

import "go-programs/RLHF/golang_alignment_5/30-10-24/105724/turn1/modelA/logging"

func main() {
	logger := logging.NewLogger()
	logger.Info("Application started.")
	logger.Warn("Something unusual happened.")
	logger.Error("Critical error, terminating.")
}
