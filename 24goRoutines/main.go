package main

import (
	"fmt"
	"time"
)

func main() {
go greeter("hello") //we command it to throw a thread but didnt mention when to came back
greeter("world")
}



func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3* time.Millisecond) //so here we tell it to show after every 5 millisec
		fmt.Println(s)
	}
}