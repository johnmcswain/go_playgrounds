package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
const omdb = "http://www.omdbapi.com/?s=Alien&i=tt3896198&apikey=5a6e37be"

func main() {
	fmt.Print(squareVal(20))
	makeAPICall()
}
func squareVal(someInt int) int{
	return someInt*someInt
}
func makeAPICall()string{
	resp, err := http.Get(omdb)
	if err == nil {
		body,err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Print(string(body))
		}
	}
	return ""
}

