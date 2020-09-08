package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	rice "github.com/GeertJohan/go.rice"
)

type Ping struct {
	Message string
}

func main() {
	appBox, err := rice.FindBox("./client/build")
	if err != nil {
		log.Fatal(err)
	}

	// static files
	http.Handle("/static/", http.FileServer(appBox.HTTPBox()))

	// SPA
	http.HandleFunc("/api/ping", apiPing)
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

func apiPing(w http.ResponseWriter, r *http.Request) {
	ping := Ping{"OK"}

	data, err := json.Marshal(ping)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
