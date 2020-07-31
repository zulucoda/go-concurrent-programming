package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Car{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		func(id int) {
			if c, ok := queryCache(id); ok {
				fmt.Println("From cache")
				fmt.Println(c)
			}
		}(id)
		func(id int) {
			if c, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				fmt.Println(c)
			}
		}(id)
		//fmt.Printf("Car not found with id: '%v'", id)
		time.Sleep(150 * time.Millisecond)
	}
	time.Sleep(2 * time.Millisecond)

}

func queryCache(id int) (Car, bool) {
	c, ok := cache[id]
	return c, ok
}

func queryDatabase(id int) (Car, bool) {
	time.Sleep(100 * time.Millisecond) // mock DB
	for _, c := range cars {
		if c.ID == id {
			cache[id] = c
			return c, true
		}
	}

	return Car{}, false
}
