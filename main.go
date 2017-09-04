package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world."))
	})
	r.HandleFunc("/.well-known/acme-challenge/{fileid}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		http.ServeFile(w, r, "./webroot/.well-known/acme-challenge/"+vars["fileid"])
	})
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8005", nil))
}
