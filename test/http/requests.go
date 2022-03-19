package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// key value arguments: https://httpbin.org/get?search=something
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln("Unable to get response", err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to get response", err)
	}
	fmt.Println(string(content))
}
