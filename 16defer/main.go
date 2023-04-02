package main

import "fmt"

func main() {
	// defer fmt.Println("one")
	// defer fmt.Println("two")
	// defer fmt.Println("three")
	// fmt.Println("Defer")
	// defer fmt.Println("in golang") //the moment we write defer it cut the line and put it to last of curly braces, which means it execute after all lines in function


	// it follows lifo order to execute
myDefer()
}

func myDefer()  {
	for i:=0; i<5; i++{
		defer fmt.Println(i)
	}
}