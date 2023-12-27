package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
	fmt.Printf("worker completed: %d\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Worker(i, &wg)
	}
	wg.Wait()
	fmt.Printf("final counter : %d\n", counter)
}
