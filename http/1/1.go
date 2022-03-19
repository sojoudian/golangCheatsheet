package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world!\n")
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println("Bytes written:", +n)
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8082", nil)
}
