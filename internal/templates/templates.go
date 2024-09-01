package templates

import (
	"embed"
	"html/template"
)

//go:embed templates
var templateDir embed.FS

func BuildTemplateCache() {

	tc := make(map[string]*template.Template)

}

func RenderTemplate() {

}
