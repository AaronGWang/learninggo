package main

import (
	"fmt"
	"time"
	"sync"
)

var mutex = sync.RWMutex{}
var wg  = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main(){
	t0  := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nResults: %v", results)
}

func dbCall(i int){
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string){
	mutex.Lock()
	results = append(results, result)
	mutex.Unlock()
}

func log(){
	mutex.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	mutex.RUnlock()
}