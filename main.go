package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, Here it is goblog!</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "<h1>About me</h1>")
	} else {
		fmt.Fprint(w, "<h1>404</h1>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
