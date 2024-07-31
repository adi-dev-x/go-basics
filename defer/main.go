package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Goroutine 1
	wg.Add(1)
	go func() {
		wg.Done()
		fmt.Println("Deferred function for Goroutine 1 executed")
		doWork("Goroutine 1", 3)
	}()

	// Goroutine 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Deferred function for Goroutine 2 executed")
		doWork("Goroutine 2", 5)
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}

func doWork(name string, seconds int) {
	fmt.Printf("%s is starting...\n", name)
	for i := 1; i <= seconds; i++ {
		fmt.Printf("%s is working... (Second %d)\n", name, i)
	}
	fmt.Printf("%s has finished.\n", name)
}
