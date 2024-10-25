package main

import (
	"log"
	"net/http"
	"os"
	"runtime/pprof"

	"github.com/blevesearch/mmap-go"
	"github.com/go-redis/redis"
)

func main() {
	// Open the data file and map it to memory.
	dataFile, err := os.Open("large_data_file.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer dataFile.Close()
	data, err := mmap.Map(dataFile, mmap.RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Unmap()

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	pprof.Lookup("heap").WriteTo(os.Stderr, 2)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Instead of creating a large struct here, use a struct with pre-allocated fields.
	//response := &Response{} // Less efficient memory allocation
	response := responsePool.Get().(*Response)
	defer responsePool.Put(response)

	// ... Rest of the code ...
	// response is filled with data
	w.Write(response.Data)
}

func getDataFromCache(key string) (string, bool) {
	// Connect to Redis server
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Get the data from the cache
	data, err := client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", false // Key not found
		}
		log.Printf("Error fetching data from cache: %v", err)
		return "", false
	}
	return data, true
}
