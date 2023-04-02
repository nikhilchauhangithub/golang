package main

import (
	"fmt"
	"io"
	"net/http"
)


const url = "https://nikhilchauhan.netlify.app/"

func main() {
	fmt.Println("web request in golang")

response,err:= http.Get(url)

if err!=nil {
	panic(err)
}

fmt.Printf("Response is of type: %T\n",response)
fmt.Println("Response is: ",response)
defer response.Body.Close() //caller should close connection everytime

databytes,err := io.ReadAll(response.Body)

if err!=nil {
	panic(err)
}
fmt.Println("Response content is: ",string(databytes))

}