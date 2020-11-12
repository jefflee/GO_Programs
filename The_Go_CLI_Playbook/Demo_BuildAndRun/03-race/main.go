package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i // i is shared, so this line fix the problem. (Variable shadowing)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}
