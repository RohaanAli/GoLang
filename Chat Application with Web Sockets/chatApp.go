package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //res,req
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

	// start the web server
	if err := http.ListenAndServe(":3000", nil); err != nil { //Hosting Port
		log.Fatal("ListenAndServe:", err)
	}

}
