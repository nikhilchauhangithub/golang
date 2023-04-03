package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web verb in golang")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
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

func PerformPostJsonRequest(){
	const myurl ="http://localhost:8000/post"
	requestBody := strings.NewReader(`
	{
		"coursename":"Lets go with golang",
		"userAge" : 21

	}
	`)

	response,err:=http.Post(myurl, "application/json", requestBody)
	if err!=nil {
		panic(err)
	}
	content,_:= ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
defer response.Body.Close()
}

func PerformPostFormRequest()  {
	const myurl ="http://localhost:8000/postform"
	
	//formdata

	data:=url.Values{}
	data.Add("firstname", "Nikhil")
	data.Add("lastname", "Chauhan")

response,err:= http.PostForm(myurl, data)
if err!=nil {
	panic(err)
}
defer response.Body.Close()
content, _:=ioutil.ReadAll(response.Body)
fmt.Println(string(content))
}