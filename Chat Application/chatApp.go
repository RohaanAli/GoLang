package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"./trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request for the template.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func main() {

	var address = flag.String("address", ":3000", "The address of the Application")
	flag.Parse()
	gomniauth.SetSecurityKey("PUT YOUR AUTH KEY HERE")
	gomniauth.WithProviders(
		facebook.New("2280011405478753", "48f54df6a418426cdc0e406e58344f66",
			"http://localhost:3000/auth/callback/facebook"),
		github.New("d22f0f061dc89517a0bb", "a642fca1fb25147fdc031b252c20205cea77f6f5",
			"http://localhost:3000/auth/callback/github"),
		google.New("799263709705-97dol722coqkcdoopo2uonm6u3cvhlaf.apps.googleusercontent.com", "zOjNIk7GMHRHHMrWzkKBgCeL",
			"http://localhost:3000/auth/callback/google"),
	)
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	go r.run()

	fmt.Println("Running the web server on port", *address)

	// start the web servers
	if err := http.ListenAndServe(*address, nil); err != nil { //Defining Port
		log.Fatal("ListenAndServe:", err)
	}
}
