package HyperText

import (
	"html/template"
	"log"
)

var (
	tpl *template.Template
)

func StartTemplateGoHtml() {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\n...\n")

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}