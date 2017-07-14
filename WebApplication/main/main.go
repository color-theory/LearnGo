package main

import (
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		w.Header().Add("Content Type", "text/html")

		pageData := page{
			"Branching Example",
			"This is my example header",
			[3]string{"Nav Item 1", "Nav Item 2", "Nav Item 3"},
		}

		templates := template.New("template")
		templates.New("test").Parse(doc)
		templates.New("headerTemplate").Parse(header)

		templates.Lookup("test").Execute(w, pageData)
	})

	http.ListenAndServe(":8000", nil)
}

type page struct{
	Title string
	Header string
	NavItems [3]string
}

const header = `
<!DOCTYPE html>
<html>
	<head>
		<title>{{.}}</title>
	</head>
`

const doc = `
	{{template "headerTemplate" .Title}}
	<body>
		<h1>{{.Header}}</h1>
		{{if eq .Title "Example Title"}}
			<p>The title matches "Example Title" so you're seeing this.</p>
		{{else}}
			<p>The title does not match "Example Title" so you're seeing the other text.</p>
		{{end}}
		<ul>
			{{range .NavItems}}
				<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
</html>
`