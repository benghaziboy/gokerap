package main

import (
    "net/http"
)

func response(rw http.ResponseWriter, request *http.Request) {
	rw.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":8080", nil)
}
