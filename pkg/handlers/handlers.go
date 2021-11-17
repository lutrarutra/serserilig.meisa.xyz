package handlers

import (
	"meisa_xyz/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.tmpl")
}
