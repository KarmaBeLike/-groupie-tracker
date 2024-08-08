package main

import (
	"log"
	"net/http"
	"os"

	"groupie/internal"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", internal.Home)
	mux.HandleFunc("/artist/", internal.ArtistPage)
	mux.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("ui"))))

	port := "8080"

	if len(os.Args) > 1 {
		port = os.Args[1]
		if port[0] == ':' {
			port = port[1:]
		}
	}

	address := "127.0.0.1:" + port
	log.Printf("Server is starting at http://%s/", address)

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
