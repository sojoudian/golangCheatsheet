package handlers

import (
	"net/http"

	"github.com/sojoudian/golangCheatsheet/http/4/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
