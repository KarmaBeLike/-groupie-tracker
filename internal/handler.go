package internal

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	err := GetData()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	err = GetRel()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	html, err := template.ParseFiles("ui/templates/index.html")
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

	if id <= 0 || id >= 53 {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	if r.URL.Path != "/"+path[1]+"/"+path[2] || err != nil {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	err = GetData()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = GetRel()
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	html, err := template.ParseFiles("ui/templates/artist.html")
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

func ErrorHandler(w http.ResponseWriter, code int) {
	errstr := ErrorPage{code, http.StatusText(code)}
	html, err := template.ParseFiles("ui/templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	err = html.Execute(w, errstr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
}
