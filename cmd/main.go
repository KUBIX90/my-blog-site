package main

import (
	"fmt"
	"net/http"

	"github.com/KUBIX90/my-blog-site.git/internal/templates"
	"github.com/a-h/templ"
)

func main() {
	component := templates.Index("hello")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
