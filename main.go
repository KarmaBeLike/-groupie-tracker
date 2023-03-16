package main

import (
	"groupie/cmd"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", cmd.Home)
	mux.HandleFunc("/artist/", cmd.ArtistPage)
	mux.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
