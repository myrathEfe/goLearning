package main

import (
	"fmt"
	"sync"
	"time"
)

var wg4 = sync.WaitGroup{}

func main() {
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg4.Add(1)
		go count()
	}
	wg4.Wait()
	fmt.Printf("\nTotal execution time : %v", time.Since(t0))
}

func count() {
	// Simulate DB call delay
	var res int
	for i := 0; i < 100000000; i++ {
		res += 1
	}
	wg4.Done()
}
