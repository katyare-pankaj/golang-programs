// package main
package main

import (
	"go-programs/RLHF/golang_random/25-11-24/360541/turn1/modelA/common"
	"go-programs/RLHF/golang_random/25-11-24/360541/turn1/modelA/consumer"
	"go-programs/RLHF/golang_random/25-11-24/360541/turn1/modelA/producer"
)

func main() {
	ch := common.MakeChannel()

	// Start the producer and consumer in goroutines
	go producer.Produce(ch)
	consumer.Consume(ch)
}
