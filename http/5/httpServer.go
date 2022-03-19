package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// handle `/` route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World! From Maziars dev machine\n")
	})

	// run server on port "9000"
	log.Fatal(http.ListenAndServe(":8080", nil))

}
