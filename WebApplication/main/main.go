package main

import (
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		w.Header().Add("Content Type", "text/html")

		pageData := page{
			"Example Title",
			"This is my example header",
		}

		tmpl, err := template.New("test").Parse(doc)
		if err == nil{
			tmpl.Execute(w, pageData)
		}
	})

	http.ListenAndServe(":8000", nil)
}

type page struct{
	Title string
	Header string
}

const doc = `
<!DOCTYPE html>
<html>
	<head>
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Header}}</h1>
	</body>
</html>
`