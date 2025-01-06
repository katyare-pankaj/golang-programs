package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"
)

type Result struct {
	Data  []byte
	Error error
}

func ReadFileWithContext(ctx context.Context, filename string) Result {
	select {
	case <-ctx.Done():
		return Result{Error: ctx.Err()}
	default:
		data, err := ioutil.ReadFile(filename)
		return Result{Data: data, Error: err}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	result := ReadFileWithContext(ctx, "example.txt")
	if result.Error != nil {
		switch err := result.Error.(type) {
		case context.DeadlineExceeded:
			fmt.Println("Error: Operation timed out")
		default:
			fmt.Println("Error reading file:", err)
		}
		return
	}

	fmt.Println(string(result.Data))
}
