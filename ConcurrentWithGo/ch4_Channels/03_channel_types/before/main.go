package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
