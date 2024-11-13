package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"strings"
)

// Example of enabling Gzip for static files in an HTTP server
func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fs.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		fs.ServeHTTP(gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
	})))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
