package main

import (
	"fmt"
	"net/http"
)

func main() {
	/* This shows you teh request for the fav icon. */
	http.HandleFunc("/", foo)
	/* comment the line below to see what's up. */
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Fprintln(w, "go look at your terminal")
}
