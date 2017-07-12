package main

import (
	"net/http"
	"strings"
	"os"
	"bufio"
)

func main() {
	http.Handle("/", new(MyHandler)) // all routes will go to a new instance of MyHandler

	http.ListenAndServe(":8000", nil) // using Go's default server multiplexer as the handler by setting it to nil
}

type MyHandler struct {
	http.Handler
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public/" + req.URL.Path
	f, err := os.Open(path)

	if err == nil { // no problems with the file read...
		bufferedReader := bufio.NewReader(f) //using this instead of a standard file-read will save tons of memory.
		var contentType string

		switch { //so far we only have html and css files so we check for those 2 and set the content type accordingly.
		case strings.HasSuffix(path, ".css"):
			contentType = "text/css" // if we don't set this, the stylesheet files will be seen as plaintext by the browser.
		case strings.HasSuffix(path, ".html"):
			contentType = "text/html"
		}
		w.Header().Add("Content-Type", contentType)
		bufferedReader.WriteTo(w)
	} else { // throw a 404 regardless of the error.
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}
