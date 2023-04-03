package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web verb in golang")
	PerformGetRequest()
}

func PerformGetRequest()  {
	const myurl = "http://localhost:8000/get"

	response,err := http.Get(myurl)

	if err!=nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)

	var responseString strings.Builder

content,_:=ioutil.ReadAll(response.Body)
byteCount,_:=responseString.Write(content)
fmt.Println(byteCount)
fmt.Println(responseString.String())
}