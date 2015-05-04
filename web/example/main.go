package main

import (
	"github.com/cyrus-and/gdb/web"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	mux, err := web.NewHandler()
	if err != nil {
		log.Fatal(err)
	}
	// static assets
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, filepath.Join("static", req.URL.Path[1:]))
	})
	// source files
	mux.HandleFunc("/source/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, req.URL.Path[len("/source"):])
	})
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
