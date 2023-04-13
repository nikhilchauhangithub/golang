package main

import (
	"fmt"
	"net/http"
	"sync"
)


var wg sync.WaitGroup //pointer

func main() {
// go greeter("hello") //we command it to throw a thread but didnt mention when to came back
// greeter("world")

websiteList:=[]string{
	"https://lco.dev",
	"https://go.dev",
	"https://gihub.com",
	"https://meta.com",

}
for _,web:=range websiteList{
	go getStatusCode(web)
	wg.Add(1)

}
wg.Wait()
}



// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3* time.Millisecond) //so here we tell it to show after every 5 millisec
// 		fmt.Println(s)
// 	}
// }


func getStatusCode(endpoint string)  {
	defer wg.Done()
	res, err :=http.Get(endpoint)
	if err!=nil {
		fmt.Println("oops in endpoint")
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}