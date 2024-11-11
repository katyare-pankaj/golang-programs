package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"github.com/influxdata/go-tsdb"
)

func main() {
	// Project ID and bucket name to store data
	projectID := "your-project-id"
	bucketName := "your-bucket-name"

	// Initialize GCS client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create GCS client: %v", err)
	}
	defer client.Close()

	// Create the bucket if it doesn't exist
	if err := createBucketIfNotExists(ctx, client, bucketName); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	// Configuring TSDB using GCS as the storage backend
	tsdbConfig := &tsdb.TSDBConfig{
		TSDBPath:         "gcs://" + bucketName, // Replace with your GCS bucket URI
		MaxSeriesPerDB:   100000,
		MaxShardDuration: 1 * time.Hour,
	}

	// Create a new TSDB instance
	tsd, err := tsdb.NewTSDB(tsdbConfig)
	if err != nil {
		log.Fatalf("Failed to create TSDB instance: %v", err)
	}
	defer tsd.Close()

	//Securely writing data to the time series
	writeSecureData(tsd)

	//Securely reading data from the time series
	readSecureData(tsd)
}

func writeSecureData(tsd *tsdb.TSDB) {
	// Write data points
	for i := 0; i < 10; i++ {
		point := tsdb.NewPoint("my_metric",
			map[string]string{"region": "us-west"},
			map[string]interface{}{"value": i},
			time.Now().Add(-time.Duration(i)*time.Minute))
		if err := tsd.WritePoint(context.Background(), point); err != nil {
			log.Fatalf("Failed to write point: %v", err)
		}
	}
}
func readSecureData(tsd *tsdb.TSDB) {
	// Read data points
	q := tsdb.NewQuery(tsdb.QuerySpec{
		Start: time.Now().Add(-2 * time.Hour),
		End:   time.Now(),
		Aggregators: []tsdb.AggregatorSpec{
			{
				Type:  "sum",
				Field: "value",
				Name:  "total_value",
			},
		},
		GroupBy: []tsdb.GroupBySpec{
			{
				Tags: []string{"region"},
			},
		},
	})
	res, err := tsd.Query(context.Background(), q)
	if err != nil {
		log.Fatalf("Failed to query TSDB: %v", err)
	}

	for _, s := range res.Series {
		fmt.Printf("Region: %s, Total Value: %f\n", s.Tags["region"], s.Values[0].Value)
	}
}

func createBucketIfNotExists(ctx context.Context, client *storage.Client, bucketName string) error {
	_, err := client.Bucket(bucketName).Attrs(ctx)
	if err != nil {
		if err == storage.ErrBucketNotExist {
			// Create the bucket if it doesn't exist
			_, err = client.CreateBucket(ctx, bucketName, nil)
			if err != nil {
				return fmt.Errorf("Failed to create bucket %q: %v", bucketName, err)
			}
		} else {
			return fmt.Errorf("Failed to get bucket %q attributes: %v", bucketName, err)
		}
	}
	return nil
}
