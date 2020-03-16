package main

import (
	"io"
	"net/http"
)

type hotdog int

/* logic for dog */
func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

/* logic for cat */
func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat
	/* This mux will handle whatever URL comes in we specify with certain logic */
	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)
	/* If you want to catch anything added onto the url, add a trailing slash, like dog does:
	//dog/this/goes/to/dog */
	http.ListenAndServe(":8080", mux)
}
