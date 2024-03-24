package handlers

import (
	"html/template"
	"net/http"
)

type test struct {
	Data string
}

func Index(w http.ResponseWriter, r *http.Request) {

	public_tmpl_files := []string{
		"templates/layout.tmpl",
		"templates/navbar.tmpl",
		"templates/index/index.tmpl",
	}

	templates := template.Must(template.ParseFiles(public_tmpl_files...))
	templates.ExecuteTemplate(w, "layout", test{Data: "Hello world"})
}
