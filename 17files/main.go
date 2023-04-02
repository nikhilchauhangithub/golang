package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("files in golang")
	content := "this is a demo content"

	file,err := os.Create("./myCustomGoFile.txt") //create file

	if err != nil{
		panic(err)
	}

length,err :=io.WriteString(file, content)
if err != nil{
	panic(err)
}
fmt.Println("length is: ",length)
defer file.Close()
}