// main.go
package main

import (
	"fmt"

	"github.com/example/cache"
	"github.com/example/utils"
)

func main() {
	data := []int{1, 2, 3, 4, 5}
	fmt.Println("Original data:", data)
	sortedData := utils.Sort(data)
	fmt.Println("Sorted data:", sortedData)
	cachedData := cache.CacheData(sortedData)
	fmt.Println("Cached data:", cachedData)
}
