package main

import (
	"net/http"
	"io/ioutil"
	"strings"
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
	data, err := ioutil.ReadFile(string(path))

	if err == nil { // no problems with the file read...
		var contentType string

		switch { //so far we only have html and css files so we check for those 2 and set the content type accordingly.
		case strings.HasSuffix(path, ".css"):
			contentType = "text/css" // if we don't set this, the stylesheet files will be seen as plaintext by the browser.
		case strings.HasSuffix(path, ".html"):
			contentType = "text/html"
		}
		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else { // throw a 404 regardless of the error.
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}
