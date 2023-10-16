package handlers

import (
	"html/template"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {

	var views = []string{
		"views/partials/base.html",
		"views/partials/navbar.html",
		"views/home/content.html",
	}

	templates := template.Must(template.ParseFiles(views...))
	templates.ExecuteTemplate(w, "base", viewInfo{ShowNavbar: true})
}
