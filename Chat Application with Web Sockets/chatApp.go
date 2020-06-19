package main

import (
	"flag"
	"fmt"
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
	temp.once.Do(func() {

		temp.templ = template.Must(template.ParseFiles(filepath.Join("templates", temp.filename)))
	})
	temp.templ.Execute(w, r)

}

func main() {

	var address = flag.String("address", ":3000", "The address of the Application")
	flag.Parse()
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/room", r)

	go r.run()

	fmt.Println("Running the web server on port", *address)

	// start the web servers
	if err := http.ListenAndServe(*address, nil); err != nil { //Defining Port
		log.Fatal("ListenAndServe:", err)
	}
}
