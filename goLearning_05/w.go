package main

import (
	"fmt"
	"sync"
	"time"
)

var wg3 = sync.WaitGroup{}
var dbData3 = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t0 := time.Now()
	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go dbCall3(i)
	}
	wg3.Wait()
	fmt.Printf("\nTotal execution time : %v", time.Since(t0))
}

func dbCall3(i int) {
	// Simulate DB call delay
	fmt.Printf(".")
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	wg3.Done()
}
