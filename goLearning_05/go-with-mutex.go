package main

import (
	"fmt"
	"sync"
	"time"
)

//still process

var m = sync.Mutex{}
var wg2 = sync.WaitGroup{}
var dbData2 = []string{"id1", "id2", "id3", "id4", "id5"}
var results2 = []string{}

func main() {
	time0 := time.Now()
	for i := 0; i < len(dbData2); i++ {
		wg2.Add(1)
		go dbCall2(i)
	}
	wg2.Wait()
	fmt.Printf("\nTotal execution time : %v", time.Since(time0))
	fmt.Printf("\nThe results are : %v", results2)
}

func dbCall2(i int) {
	//Simulating DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is :", dbData2[i])
	m.Lock()
	results2 = append(results2, dbData2[i])
	m.Unlock()
	wg2.Done()
}

func save(result string) {
	m.Lock()
	results2 = append(results2, result)
	m.Unlock()
}
