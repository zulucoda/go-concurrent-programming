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
	go func(ch <-chan int, wg *sync.WaitGroup) {
		if msg, ok := <-ch; ok {

			fmt.Println(msg, ok)
		}
		wg.Done()
	}(ch, wg)

	// send channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(ch)
		wg.Done()

	}(ch, wg)
	wg.Wait()
}
