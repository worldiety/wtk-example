package main

import (
	"flag"
	"fmt"
	"github.com/lpar/gzipped/v2"
	"github.com/worldiety/wtk/theme/material"
	"log"
	"net/http"
	"path"
	"strings"
)

func withIndexHTML(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			fmt.Println(r.URL.Path)
			newpath := path.Join(r.URL.Path, "index.html")
			r.URL.Path = newpath
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", ".", "the directory of static files to host")
	flag.Parse()

	material.Resources(http.DefaultServeMux)
	http.Handle("/", withIndexHTML(gzipped.FileServer(gzipped.Dir(*directory))))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}