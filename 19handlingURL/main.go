package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("handling URLs in golang")

	const myurl string ="https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

	// fmt.Println(myurl)

	//parsing

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)


	qParams := result.Query()
	fmt.Printf("the type of query params are: %T\n", qParams)
	fmt.Println("the type of query params are:", qParams)
	fmt.Println(qParams["v"])

	for _, val := range qParams {
fmt.Println("Param is:", val)
	}


	//construct url

	partsOfURL :=&url.URL{
		Scheme: "https",
		Host: "www.youtube.com",
		Path: "/watch",
		RawPath: "v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26",
	}

	mainURL := partsOfURL.String()
	fmt.Println(mainURL)
}