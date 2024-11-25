// package consumer
package consumer

import (
	"fmt"
	"go-programs/RLHF/golang_random/25-11-24/360541/turn1/modelA/common"
	"time"
)

// Consume receives messages from the channel and processes them.
func Consume(ch common.ChannelType) {
	for message := range ch {
		fmt.Println("Consumer receiving:", message)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Consumer done.")
}
