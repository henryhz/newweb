package main

import (
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello world")
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil)
}