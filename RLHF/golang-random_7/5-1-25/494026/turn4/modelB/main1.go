package main

import "fmt"

func sieveOfEratosthenes(n int, c chan<- int) {
	if n < 2 {
		return
	}

	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	isPrime[0] = false
	isPrime[1] = false

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			c <- i
		}
	}
	close(c)
}

func main() {
	n := 30
	c := make(chan int)

	go sieveOfEratosthenes(n, c)

	for prime := range c {
		fmt.Println(prime)
	}
}
