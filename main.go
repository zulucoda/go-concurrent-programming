package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Car{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	// creating a pointer to the wait group because we going to passing around
	// the waitGroup in the application.
	// Note: you dont want to copy a waitGroup as you pass it around. Thats why you use a
	// pointer.
	wg := &sync.WaitGroup{}

	// mutex is similar to a waitGroup, used to handle race conditions
	// Note: you dont want to copy a mutex as you pass it around. Therefore use a pointer
	mutex := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		// This lets the waitGroup know that there 2 task that we are waiting on
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
			if c, ok := queryCache(id, mutex); ok {
				fmt.Println("From cache")
				fmt.Println(c)
			}
			// When the task is done, we let the waitGroup know
			wg.Done()
		}(id, wg, mutex)
		go func(id int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
			if c, ok := queryDatabase(id, mutex); ok {
				fmt.Println("from database")
				fmt.Println(c)
			}
			// When the task is done, we let the waitGroup know
			wg.Done()
		}(id, wg, mutex)
		//fmt.Printf("Car not found with id: '%v'", id)
		//time.Sleep(150 * time.Millisecond)
	}

	// This wait, will wait for all tasks to be done.
	wg.Wait()

}

func queryCache(id int, mutex *sync.RWMutex) (Car, bool) {

	// This allow mutliple goroutines to read from the cache
	mutex.RLock()
	c, ok := cache[id]
	mutex.RUnlock()
	return c, ok
}

func queryDatabase(id int, mutex *sync.RWMutex) (Car, bool) {
	time.Sleep(100 * time.Millisecond) // mock DB
	for _, c := range cars {
		if c.ID == id {

			// whichever goroutine called this, now owns the mutex. So nothing else is going to be able
			// to access procteded code until the owing goroutine calls UnLock
			mutex.Lock()
			cache[id] = c
			mutex.Unlock()
			return c, true
		}
	}

	return Car{}, false
}
