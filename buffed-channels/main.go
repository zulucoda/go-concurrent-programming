package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("vim-go")

	wg := &sync.WaitGroup{}

	// creating a buffered channel with 1 message
	ch := make(chan int, 1)

	wg.Add(2)

	// receiving channel
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	// send channel
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		ch <- 27
		wg.Done()

	}(ch, wg)
	wg.Wait()
}
