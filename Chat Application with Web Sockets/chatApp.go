package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request for the template.
func (temp *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if temp.templ == nil {
		temp.templ = template.Must(template.ParseFiles(filepath.Join("templates", temp.filename)))
	}
	temp.templ.Execute(w, nil)

}

func main() {

	http.Handle("/", &templateHandler{filename: "/index.html"})

	// start the web servers
	if err := http.ListenAndServe(":3000", nil); err != nil { //Defining Port
		log.Fatal("ListenAndServe:", err)
	}
}
