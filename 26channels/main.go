package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang")

	mych:=make(chan int,2)
	wg:=&sync.WaitGroup{}

	wg.Add(2)

	go func(ch <- chan int, wg *sync.WaitGroup) { //receive only
		val,isChannelOpen:= <-mych

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(mych,wg)

	go func(ch chan <- int, wg *sync.WaitGroup) { //send only
		
mych<-5
		close(mych)
		
		wg.Done()
	}(mych,wg)
	wg.Wait()
}