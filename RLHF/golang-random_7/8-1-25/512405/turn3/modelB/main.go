func consumer(id int, dataCh <-chan int, stopCh <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case randInt, ok := <-dataCh:
			if !ok {
				// dataCh has been closed, exit the loop
				fmt.Printf("Consumer %d: Data channel closed, exiting\n", id)
				return
			}
			fmt.Printf("Consumer %d: Consumed item %d\n", id, randInt)
			time.Sleep(consumerSleep)
		case <-stopCh:
			// Consumer received a stop signal, exit gracefully
			fmt.Printf("Consumer %d: Received stop signal, exiting\n", id)
			return
		default:
			// No data available, do some other work or sleep before checking for stop signal
			time.Sleep(10 * time.Millisecond)
		}
	}
}
