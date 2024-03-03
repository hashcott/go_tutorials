package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6", "id7", "id8", "id9", "id10"}
var wg = sync.WaitGroup{}
var m = sync.Mutex{}
var rm = sync.RWMutex{}

var results []string = []string{}

func main() {
	t0 := time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\n Total execution time: %v\n", time.Since(t0))
	fmt.Println(results)
}

func dbCall(i int) {
	// simulate a database call
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result of the database call is: ", dbData[i])
	save(dbData[i])
	log()

	wg.Done()
}

func save(data string) {
	m.Lock() // lock the resource
	results = append(results, data)
	m.Unlock() // unlock the resource, so other goroutines can access it
}

func log() {

	rm.RLock() // lock the resource, but allow multiple readers, but no writers
	fmt.Println("Current value: ", results)
	rm.RUnlock()
}
