package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient
	//resp, err := http.Post("https://httpbin.org/post", "text/plain",
	//	strings.NewReader("This is the request content"))
	req, err := http.NewRequest("GET", "http://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln("Unable to get response", err)
	}
	resp, err := client.Do(req)
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to get response", err)
	}
	fmt.Println(string(content))
}
