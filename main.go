package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Host struct {
	Host string
	Date string
}

func main() {
	// use PORT environment variable, or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	tmpl := template.Must(template.ParseFiles("layout.html"))
	fs := http.FileServer(http.Dir("static/"))
	server.Handle("/static/", http.StripPrefix("/static/", fs))
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving request: %s", r.URL.Path)
		host, _ := os.Hostname()
		date := time.Now().Local().Format("2006-01-02")
		data := Host{
			Host: host,
			Date: date,
		}
		tmpl.Execute(w, data)
	})

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

