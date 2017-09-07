package main

import "goWorkspace/goforsre/cms"
import "os"

func main() {
	p := &cms.Page{
		Title:   "Hello, World!",
		Content: "This is the body of out webpage",
	}
	cms.Temp.ExecuteTemplate(os.Stdout, "index", p)
}
