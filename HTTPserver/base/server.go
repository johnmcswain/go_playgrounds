package main

import (
"io"
"net/http"
)

func responseMessage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "The go Server has received a "+r.Method+" request")
}

func main() {
	http.HandleFunc("/", responseMessage)
	http.ListenAndServe(":8001", nil)
}
