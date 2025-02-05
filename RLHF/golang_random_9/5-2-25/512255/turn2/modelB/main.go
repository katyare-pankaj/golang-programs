package main  

import (  
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "math/rand"
        "net/http"
        "sync"
        "time"
)

const (  
        port             = ":8080"         // Add your microservice port
        numRequests       = 1000           // Total number of requests to send
        concurrency       = 100            // Concurrency level
        requestDuration  = 5 * time.Second // Duration of each request
        maxBodySize      = 1024           // Maximum size of request body in bytes
)

type RequestReport struct {  
        ResponseTime time.Duration   `json:"response_time"`
        Code         int              `json:"code"`
}

type Statistics struct {  
        Total         int             `json:"total"`
        Min           time.Duration   `json:"min"`
        Max           time.Duration   `json:"max"`
        Average       time.Duration   `json:"average"`
        Count200      int             `json:"count_200"`
        Count500      int             `json:"count_500"`
        CountOther    int             `json:"count_other"`
}

func main() {  
        //Create a channel to receive reports from workers
        reports := make(chan RequestReport, numRequests)

        //Start a goroutine to process reports
        go processReports(reports)

        wg := &sync.WaitGroup{}
        wg.Add(numRequests)

        startTime := time.Now()

        // Create and send concurrent requests
        for i := 0; i < numRequests; i++ {  
                go sendRequest(wg, reports)
        }

        wg.Wait()

        endTime := time.Now()
        totalTime := endTime.Sub(startTime)

        fmt.Printf("Total time for %d requests with concurrency %d: %v\n", numRequests, concurrency, totalTime)
}  

func sendRequest(wg *sync.WaitGroup, reports chan<- RequestReport) {  
        defer wg.Done()

        // Generate random request body
        body := generateRandomRequestBody(maxBodySize)

        client := &http.Client{Timeout: requestDuration}
        req, err := http.NewRequest("POST", "http://localhost"+port+"/benchmark", body)
        if err != nil {  
                log.Printf("Error creating request: %v\n", err)
                return
        }

        start := time.Now()
        resp, err := client.Do(req)
        if err != nil {  
                log.Printf("Request failed: %v\n", err)
                reports <- RequestReport{ResponseTime: time.Since(start), Code: 0}
                return
        }
        defer resp.Body.Close()

        duration := time.Since(start)
        reports <- RequestReport{ResponseTime: duration, Code: resp.StatusCode}
}  

func generateRandomRequestBody(maxSize int) []byte {  
        bodySize := rand.Intn(maxSize) + 1
        body := make([]byte, bodySize)
        rand.Read(body)
        return body
}  

func processReports(reports <-chan RequestReport) {  
        var totalTime time.Duration
        var totalRequests int
        minTime := time.Duration(1 << 63) - 1 // Max duration value
        maxTime := time.Duration(0)
        count200 := 0
        count500 := 0
        var otherCodes map[int]int = make(map[int]int)

        for report := range reports {  
                totalTime += report.ResponseTime