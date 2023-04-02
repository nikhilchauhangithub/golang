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

// call readFile

readFile("./myCustomGoFile.txt")
}

func readFile(fileName string)  {
	dataByte,err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataByte))
}