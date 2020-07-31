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

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		// This lets the waitGroup know that there 2 task that we are waiting on
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup) {
			if c, ok := queryCache(id); ok {
				fmt.Println("From cache")
				fmt.Println(c)
			}
			// When the task is done, we let the waitGroup know
			wg.Done()
		}(id, wg)
		go func(id int, wg *sync.WaitGroup) {
			if c, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				fmt.Println(c)
			}
			// When the task is done, we let the waitGroup know
			wg.Done()
		}(id, wg)
		//fmt.Printf("Car not found with id: '%v'", id)
		//time.Sleep(150 * time.Millisecond)
	}

	// This wait, will wait for all tasks to be done.
	wg.Wait()

}

func queryCache(id int) (Car, bool) {
	c, ok := cache[id]
	return c, ok
}

func queryDatabase(id int) (Car, bool) {
	time.Sleep(100 * time.Millisecond) // mock DB
	for _, c := range cars {
		if c.ID == id {
			//	cache[id] = c
			return c, true
		}
	}

	return Car{}, false
}
