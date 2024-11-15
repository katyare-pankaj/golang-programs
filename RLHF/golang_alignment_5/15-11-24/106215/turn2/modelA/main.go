package main

import (
	"fmt"
	"runtime"
	"time"
)

var total int

func naiveFactorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * naiveFactorial(n-1)
}
func iterativeFactorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
func primeSieve(n int) ([]int, bool) {
	isPrime := make([]bool, n+1)
	primes := []int{}

	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return primes, true
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	startTime := time.Now()
	func1()
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Execution duration of functin1:%s \n", duration)

	startTime = time.Now()
	func2()
	endTime = time.Now()
	duration = endTime.Sub(startTime)
	fmt.Printf("Execution duration of func2:%s\n", duration)

}

func func1() {

	/* This is to  See the common tail call optimization ; in goc compilation of default of this is 1}*,Hence,  no real impact is seen directly here.*/

	for i := 0; i < 10000; i++ {
		fmt.Println(naiveFactorial(15))
		total++
	}
}

func func2() {
	/*argument starts at  i from the local higher function,identifier seekar into the samefunction; hence SROA suffers here, but goo recovery wush loo read
	 */

	primes, okay := primeSieve(100000)
	if !okay {
		panic("Failed to sieve primes")
	}
	for i := range primes {
		total += iterativeFactorial(primes[i])
	}

}
