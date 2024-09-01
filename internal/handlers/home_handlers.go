package handlers

import (
	"html/template"
	"net/http"
)

func HandleGetHome(w http.ResponseWriter, r *http.Request) {
	html := `<html><head><title>Home Page</title></head><body><h1>Hello {{ .name }}</h1></body></html>`

	tmpl, err := template.New("home").Parse(html)

	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})

	data["name"] = "Andrew"

	err = tmpl.Execute(w, data)

	if err != nil {
		panic(err)
	}

}
