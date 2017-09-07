package cms

import "html/template"

var Temp = template.Must(template.ParseGlob("../templates/*"))

type Page struct {
	Title string
	Content string

}

