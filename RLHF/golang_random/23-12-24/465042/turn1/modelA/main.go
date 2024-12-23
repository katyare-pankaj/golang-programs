package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// BookFlight simulates booking a flight.
func BookFlight(wg *sync.WaitGroup, itineraryName string) {
	defer wg.Done()
	fmt.Printf("Booking flight for %s...\n", itineraryName)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Printf("Flight for %s booked successfully!\n", itineraryName)
}

// BookHotel simulates booking a hotel.
func BookHotel(wg *sync.WaitGroup, itineraryName string) {
	defer wg.Done()
	fmt.Printf("Booking hotel for %s...\n", itineraryName)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Printf("Hotel for %s booked successfully!\n", itineraryName)
}

// BookRentalCar simulates booking a rental car.
func BookRentalCar(wg *sync.WaitGroup, itineraryName string) {
	defer wg.Done()
	fmt.Printf("Booking rental car for %s...\n", itineraryName)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Printf("Rental car for %s booked successfully!\n", itineraryName)
}

func main() {
	var wg sync.WaitGroup

	itineraries := []string{"New York Trip", "Paris Getaway", "Tokyo Adventure"}

	for _, itineraryName := range itineraries {
		wg.Add(3) // Add 1 for each concurrent task

		go BookFlight(&wg, itineraryName)
		go BookHotel(&wg, itineraryName)
		go BookRentalCar(&wg, itineraryName)
	}

	fmt.Println("Waiting for all itinerary tasks to complete...")
	wg.Wait()

	fmt.Println("All itinerary tasks completed.")
}
