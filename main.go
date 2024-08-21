package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fs := http.FileServer(http.Dir("./static/files"))
	http.Handle("/files/", http.StripPrefix("/files/", fs))
	http.ListenAndServe(":8080", nil)
}
