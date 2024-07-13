package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

var wg = sync.WaitGroup{} // WaitGroup just counters.
// var m = sync.Mutex{}    // Use Lock method and Unlock method to protect the code.
var m = sync.RWMutex{} // Use RLock method and RUnlock method to protect the code.
var results []string

func main() {
	t0 := time.Now()

	for i := 0; i < len(dbData); i++ {
		/*
			When we spawn a new go routine we
			make sure to add one to the counter.
			And we will call the done method
			at the end inside the function we used.
		*/
		wg.Add(1)
		go dbCall(i)
	}
	/*
		We call the weight method this method is going to
		wait for the counter to go back down to zero.
	*/
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Printf("The results: %v\n", results)
}

func dbCall(i int) {
	/*
		Actually we need to mutex to control the writing to our slice.
		We should care about corrupting memory.

		But there is still a question that the mutex may destroy the concurrency of database calls.
		To solve this problem, we can use RWMutex.
	*/
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	//fmt.Println("The result from the database is ", dbData[i])
	//m.Lock() // It'll wait for another go routine release the lock.
	//results = append(results, dbData[i])
	//m.Unlock()

	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("The current results are: %v\n", results)
	m.RUnlock()
}
