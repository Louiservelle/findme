package models

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
	Templates renderer we parsefiles with the template name and add .page.tmpl extension,
	to render the good tempale.

	and Execute templates.
*/

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		fmt.Println("error parsing template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
