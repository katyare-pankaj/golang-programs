// Example of simple data sharding in Go
package main

import (
	"fmt"
	"hash/fnv"
)

func getShardNumber(key string, numShards int) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() % uint32(numShards))
}

func main() {
	logDataKey := "some_log_data_key"
	numShards := 10
	shardNumber := getShardNumber(logDataKey, numShards)
	fmt.Println("Shard Number for", logDataKey, ":", shardNumber)
	// Shard number will be between 0 and numShards-1
}
