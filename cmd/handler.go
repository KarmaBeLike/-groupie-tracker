package cmd

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, 404)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	GetData()
	GetRel()
	html, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = html.Execute(w, artists)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(path[2])
	if r.URL.Path != "/"+path[1]+"/"+path[2] || err != nil {
		ErrorHandler(w, 404)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	err = GetData()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
	}
	err = GetRel()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
	}
	html, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = html.Execute(w, artists[id-1])
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

type ErrorPage struct {
	Code    int
	Message string
}

func ErrorHandler(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	errstr := ErrorPage{code, http.StatusText(code)}
	html, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	err = html.Execute(w, errstr)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
}
