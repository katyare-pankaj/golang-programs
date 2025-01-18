package main

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

// A sample third-party library function that might not handle defer correctly.
func thirdPartyFunction() error {
	defer fmt.Println("Defer from thirdPartyFunction.")
	return errors.New("something went wrong")
}

// Wrap the third-party function to handle defer correctly.
func callThirdPartyFunction() error {
	defer fmt.Println("Defer from callThirdPartyFunction.")
	return thirdPartyFunction()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Call the wrapped third-party function
		err := callThirdPartyFunction()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}()
	wg.Wait()
	fmt.Println("Main function finished.")
}
