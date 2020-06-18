package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //res,req on root page
		w.Write([]byte(`
      <html>
        <head>
          <title>Chat Room</title>
        </head>
        <body>
          Let's chat!
        </body>
      </html>
    `))
	})

	// start the web servers
	if err := http.ListenAndServe(":3000", nil); err != nil { //Defining Port
		log.Fatal("ListenAndServe:", err)
	}
}
