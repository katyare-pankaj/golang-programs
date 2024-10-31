package main

import (
	"fmt"
	"time"

	"github.com/cucumber/godog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User represents a user in the system
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100)"`
	Email     string `gorm:"type:varchar(255)"`
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Synchronizer handles data synchronization
type Synchronizer struct {
	db *gorm.DB
}

// NewSynchronizer creates a new Synchronizer
func NewSynchronizer() *Synchronizer {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	return &Synchronizer{db: db}
}

// SynchronizeUsers synchronizes users from one database to another
func (s *Synchronizer) SynchronizeUsers() error {
	// Implement your synchronization logic here
	// For this example, we'll just perform a simple two-way synchronization
	return nil
}

func FeatureContext(s *godog.ScenarioContext) {
	var synchronizer *Synchronizer

	s.BeforeScenario(func(*godog.Scenario) {
		synchronizer = NewSynchronizer()
	})

	s.Step(`^I synchronize users$`, func() error {
		return synchronizer.SynchronizeUsers()
	})

	s.Step(`^the users should be synchronized successfully$`, func() error {
		// Add your assertion logic here to verify that the users are synchronized correctly
		return nil
	})
}

func main() {
	status, err := godog.TestSuites(
		&godog.TestSuite{
			Name:                "Synchronization features",
			ScenarioInitializer: FeatureContext,
			FeaturesPath:        "features",
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	godog.Exit(status)
}
