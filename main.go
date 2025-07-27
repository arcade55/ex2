package main

import (
	"ex1/pages"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//fmt.Println(pages.Home().Render())

	// 1. Create a new ServeMux (router).
	mux := http.NewServeMux()

	// 2. Register handlers on the new mux.
	// This handler now only accepts GET requests for the root path.
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, pages.Home().Render())
	})

	fileServer := http.FileServer(http.Dir("static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	// 3. Start the server, passing your new mux instead of nil.
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
