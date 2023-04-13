package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition in golang")

	wg := &sync.WaitGroup{}
	mut:=&sync.RWMutex{}
	var score =[]int{0}
wg.Add(3) //tell, how many goroutines are there
	go func(wg *sync.WaitGroup, m *sync.RWMutex)  {
		fmt.Println("One")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg,mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex)  {
		fmt.Println("Two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg,mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex)  {
		fmt.Println("Three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	wg.Wait()
	fmt.Println(score)
}