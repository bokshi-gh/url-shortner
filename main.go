package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"url-shortner/handlers"
)

func main() {
	http.HandleFunc("/register", handlers.RegisterHandler)
	// http.HandleFunc("/login", handlers.LoginHandler)
	// http.HandleFunc("/delete-account", handlers.DeleteAccountHandler)

	// http.HandleFunc("/add-url", handlers.AddURLHandler)
	// http.HandleFunc("/delete-url", handlers.DeleteURLHandler)
	// http.HandleFunc("/r/", handlers.RedirectURL)


	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("./static", r.URL.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// If file doesn't exist, serve index.html (SPA support)
			http.ServeFile(w, r, "./static/index.html")
			return
		}
		fs.ServeHTTP(w, r)
	}))

	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

