package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	//client := http.
	resp, err := http.Post("https://httpbin.org/post", "text/plain",
		strings.NewReader("This is the request content"))
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
