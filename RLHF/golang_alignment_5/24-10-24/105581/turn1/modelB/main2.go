// main.go

package main

import (
	"fmt"

	"github.com/example/campaignmanager"
	"github.com/example/datasource"
)

func main() {
	// Create an in-memory data source
	dataSource := datasource.NewInMemoryDataSource()

	// Create a campaign manager using the in-memory data source
	campaignManager := campaignmanager.NewCampaignManager(dataSource)

	// Start real-time data syncing
	campaignManager.StartSyncing()

	// Add some subscribers
	campaignManager.AddSubscriber("alice@example.com", map[string]string{"name": "Alice", "age": "25"})
	campaignManager.AddSubscriber("bob@example.com", map[string]string{"name": "Bob", "age": "30"})

	// Get the current list of subscribers
	subscribers := campaignManager.GetSubscribers()
	fmt.Println("Current Subscribers:")
	for email, data := range subscribers {
		fmt.Printf("%s: %v\n", email, data)
	}
}
