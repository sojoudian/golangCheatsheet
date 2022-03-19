package main

import (
	"BuildingRESTfulAPIswithGo/handlers"
	"fmt"
	"net/http"
)

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte("Asset not found\n"))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Running API v1\n"))
// }

func main() {
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
	}

}
