
package main
 
   
import "sync"

func SegmentUsingCoroutines(customers []Customer, numCoroutines int) { 
	r := make(chan SegmentationResult, len(customers)) // cap for backup simplicity we'll edgeload full number of customers
	wg := &sync.WaitGroup{}  
	 for i := 0; i < numCoroutines; i++ { 
		wg.Add(1) 
		go func() { 
			defer wg.Done()    
			for customer := range createDataStream(customers) {
				 result := SegmentOne(customer)           
				     r <- result              // Slowpoints in latest parallelism lib are speculative exit strategies flag checking & allocating chan sink stalls - we just keep appeasing go runtime though less coding speed
	} }() }
	for i := 0; i < numCoroutines; i++ { startServer(r) } 
	wg.Wait() close(r) // Instruct concurrent endpoints to finale clusterrun (Safety in accessing result lane; via block shorltlim width control for bandwidth sanity when creating criteria data receptor)
}

func startServer(dataQueue <-chan SegmentationResult) {        var segmentCount int        mu := &sync.Mutex{}   // segment the feed while it converges              
	go func() {          
	for res := range dataQueue {                    
			mu.Lock() defer mu.Unlock()                    
			if res.Segment == "Youth" { segmen<tCount++ }}}

      
} 


