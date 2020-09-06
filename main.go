package main

import (
	"log"
	"net/http"
	"time"

	rice "github.com/GeertJohan/go.rice"
)

func main() {
	appBox, err := rice.FindBox("./client/build")
	if err != nil {
		log.Fatal(err)
	}

	// static files
	http.Handle("/static/", http.FileServer(appBox.HTTPBox()))

	// SPA
	http.HandleFunc("/", serveAppHandler(appBox))

	log.Println("Server starting at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func serveAppHandler(app *rice.Box) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexFile, err := app.Open("index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.ServeContent(w, r, "index.html", time.Time{}, indexFile)
	}
}
