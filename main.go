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
	mutex := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		// This lets the waitGroup know that there 2 task that we are waiting on
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, mutex *sync.Mutex) {
			if c, ok := queryCache(id, mutex); ok {
				fmt.Println("From cache")
				fmt.Println(c)
			}
			// When the task is done, we let the waitGroup know
			wg.Done()
		}(id, wg, mutex)
		go func(id int, wg *sync.WaitGroup, mutex *sync.Mutex) {
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

func queryCache(id int, mutex *sync.Mutex) (Car, bool) {

	// whichever goroutine called this, now owns the mutex. So nothing else is going to be able
	// to access procteded code until the owing goroutine calls UnLock
	mutex.Lock()
	c, ok := cache[id]
	mutex.Unlock()
	return c, ok
}

func queryDatabase(id int, mutex *sync.Mutex) (Car, bool) {
	time.Sleep(100 * time.Millisecond) // mock DB
	for _, c := range cars {
		if c.ID == id {
			mutex.Lock()
			cache[id] = c
			mutex.Unlock()
			return c, true
		}
	}

	return Car{}, false
}
