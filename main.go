package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/{path}", anyHandler)
	r.HandleFunc("/", anyHandler)
	http.Handle("/", loghttp(r))
	log.Println("Listening on 8085...")
	http.ListenAndServe(":8085", nil)
}

// loghttp just logs the path and headers
func loghttp(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, r.Header)
		h.ServeHTTP(w, r)
	})
}

func anyHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	headers := r.Header
	query := r.URL.Query()
	log.Printf("%+v %+v %+v", params, headers, query)

	var echo = struct {
		Headers http.Header       `json:"headers"`
		Query   url.Values        `json:"query"`
		Params  map[string]string `json:"params"`
		Method  string            `json:"method"`
	}{
		r.Header,
		r.URL.Query(),
		mux.Vars(r),
		r.Method,
	}

	echodata, _ := json.MarshalIndent(&echo, "", "  ")

	w.Write(echodata)
}
