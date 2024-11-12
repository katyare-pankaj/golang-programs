
package main  
import (  
    "fmt"
    "sync"
    "time"
    "github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

const (
    // Replace these values with your actual Aliyun Table Store credentials
    accessKeyId     = "your_access_key_id"
    accessKeySecret = "your_access_key_secret"
    instanceName    = "your_instance_name"
    tableName       = "rate_limiting_table"
    primaryKeyName  = "api_name"
)

type RateLimiter struct {
    mu         sync.Mutex
    rateLimits map[string]rateLimit
    client     *tablestore.TableStoreClient
}

type rateLimit struct {
    requestsPerMinute int
    lastRequestTime   time.Time
    remainingRequests int
}

func NewRateLimiter() (*RateLimiter, error) {
    client, err := tablestore.NewClient(accessKeyId, accessKeySecret, instanceName)
    if err != nil {
        return nil, fmt.Errorf("error creating Table Store client: %w", err)
    }
    
    rl := &RateLimiter{
        rateLimits: make(map[string]rateLimit),
        client:     client,
    }
    
    // Initialize the rate limits from the Table Store
    if err := rl.initializeRateLimits(); err != nil {
        return nil, fmt.Errorf("error initializing rate limits: %w", err)
    }
    
    return rl, nil
}

func (rl *RateLimiter) initializeRateLimits() error {
    // Replace this with your actual Table Store query to fetch rate limits
    query := tablestore.NewGetRowQuery()
    query.SetTableName(tableName)
    query.SetPrimaryKey([]tablestore.PrimaryKeyColumn{
        {
            Name: primaryKeyName,
            Value: tablestore.PrimaryKeyValue{
                Type:  tablestore.PrimaryKeyType_STRING,
                Value: "your_api_name", // Replace with the actual API name
            },
        },
    })
    
    resp, err := rl.client.GetRow(query)
    if err != nil {
        return fmt.Errorf("error fetching rate limits from Table Store: %w", err)
    }
    
    if !resp.IsOk() {
        return fmt.Errorf("error fetching rate limits: %s", resp.GetError())
    }
    
    // Parse the rate limits from the response and update the internal map
    // ... (Implementation details omitted)
    
    return nil
}

func (rl *RateLimiter) GetRateLimit(apiName string) (int, error) {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    lim, ok := rl.rateLimits[apiName]
