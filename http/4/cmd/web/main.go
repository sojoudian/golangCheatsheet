package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sojoudian/golangCheatsheet/http/4/pkg/config"
	"github.com/sojoudian/golangCheatsheet/http/4/pkg/handlers"
	"github.com/sojoudian/golangCheatsheet/http/4/render"
)

const portNumber = ":8082"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache: ", err)
	}
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
