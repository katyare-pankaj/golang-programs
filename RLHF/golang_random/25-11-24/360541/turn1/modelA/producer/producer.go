// package producer
package producer

import (
	"fmt"
	"go-programs/RLHF/golang_random/25-11-24/360541/turn1/modelA/common"
	"time"
)

// Produce generates messages and sends them over the channel.
func Produce(ch common.ChannelType) {
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Message %d", i)
		fmt.Println("Producer sending:", message)
		ch <- message
		time.Sleep(1 * time.Second)
	}
	close(ch)
}
